package generator

import (
	"fmt"
	"math/rand/v2"

	// Needed to fake personal data
	"github.com/go-faker/faker/v4"
)

// The main helper function to be exported out of github.com/jamesthecog/generator.  This
// function fakes data in the format expected of batch and single data - it's also
// the only thing that's exported in this function (i.e., everything else are private functions
// that belong to a part of the faked data).
//
// The data being faked only contains four sections for now: EQ5D5L responses, patient information,
// MUST questionnaire items, and distress thermometer items:
func FakeData() (map[string]any, error) {
	fakeData := make(map[string]any)
	personalInfo := fakePersonalInfo()
	dtInfo, err := fakeDistressInfo()
	if err != nil {
		return nil, err
	}
	mustInfo, err := fakeMUSTData(personalInfo["gender"])
	if err != nil {
		return nil, err
	}
	eq5d5lInfo, err := fakeEQ5D5LData()
	if err != nil {
		return nil, err
	}

	fakeData["patient_info"] = personalInfo
	fakeData["distress_thermometer"] = dtInfo
	fakeData["must"] = mustInfo
	fakeData["eq5d5l"] = eq5d5lInfo
	return fakeData, nil
}

// Constants - mainly file paths from the root directory to read in the JSON files
// contained within ./data/:
const (
	// Personal info. constants:
	GenderThres = 0.5 // --> Threshold for gender

	// Distress Thermometer constants:
	distressDataPath = "./data/dt.json" // --> Path to dt.json
	maxDistress      = 8                // --> The maximum distress thermometer score
	fallbackAmt      = 3                // --> Fallback max. amount of issues to select per category
	badDTThres       = 0.3              // --> Threshold for questionable data

	// MUST data
	minMaleWeight   = 51   // --> Minimum allowed male weight
	maxMaleWeight   = 100  // --> Maximum allowed male weight
	minFemaleWeight = 44   // --> Ditto (but for women)
	maxFemaleWeight = 86   // --> Ditto (but for women)
	minMaleHeight   = 1.7  // --> Minimum allowed male height
	maxMaleHeight   = 1.82 // --> Maximum allowed male height
	minFemaleHeight = 1.55 // --> Ditto (but for women)
	maxFemaleHeight = 1.65 // --> Ditto (but for women)
	badMUSTThres    = 0.3  // --> Threshold for questionable data

	// For EQ5D5L data:
	eq5d5lPath      = "./data/eq5d5l.json" // --> Path to possible EQ5D5L responses
	eq5d5lRatingMax = 101                  // --> The maximum rating of the questionnaire's rating response
	eq5d5lThres     = 0.15                 // --> The threshold for a malformed rating
)

// -- Private helper functions --

// Fakes personal information with faker's personal information API.  A random float
// will be used to determine if the person in question will be a male or a female.
//
// If rand.Float64() is smaller than or equal to GenderThres, then let our patient be a male and
// vice versa.
func fakePersonalInfo() map[string]string {
	randFloat := rand.Float64()
	personalInfo := make(map[string]string)

	if randFloat <= GenderThres {
		personalInfo["name"] = fmt.Sprintf("%s %s %s", faker.TitleFemale(), faker.FirstNameFemale(), faker.LastName())
		personalInfo["gender"] = "female"
	} else {
		personalInfo["name"] = fmt.Sprintf("%s %s %s", faker.TitleMale(), faker.FirstNameMale(), faker.LastName())
		personalInfo["gender"] = "male"
	}

	address, country := faker.GetRealAddress(), faker.GetCountryInfo()
	personalInfo["home_address"] = address.Address
	personalInfo["nationality"] = country.Name
	return personalInfo
}

// Fakes distress thermometer information.  This generator depends on the initially
// generated distress score to select random items from the list; the rand.Float64()
// function will be used to determine whether an appropriate item will be chosen for
// each distress category.
//
// I've also written this function such that if rand.Float64() is lesser than
// RightWrongThres, then we'll return a potentially bad slice.
func fakeDistressInfo() (map[string]any, error) {
	dtData, err := loadJson[map[string][]string](distressDataPath)
	if err != nil {
		return nil, fmt.Errorf("Failed to load in DT data: %v", err)
	}

	result, dtScore := make(map[string]any), rand.IntN(maxDistress)
	result["distress_rating"] = dtScore
	categories := []string{"emotional", "family", "practical", "physical"}

	for k, v := range dtData {
		var toSample []string
		numProblems := rand.IntN(max(fallbackAmt, dtScore))
		if rand.Float64() < badDTThres {
			trimmedCats := removeElement(categories, k)
			badCat := trimmedCats[rand.IntN(len(trimmedCats))]
			toSample = append(v, dtData[badCat]...)
		} else {
			toSample = v
		}
		res, err := randomSample(toSample, numProblems)
		if err != nil {
			return nil, err
		}
		result[k] = res
	}
	return result, nil
}

// Generates synthetic MUST data - it's a little verbose, but
// there's at least some degree of reality to the numbers (i.e., I
// gathered them off Google)!
func fakeMUSTData(gender string) (map[string]any, error) {
	result := make(map[string]any)
	choices := []string{"yes", "no"}
	var beforeWeight, curWeight, height float64

	if gender == "male" {
		if rand.Float64() < badMUSTThres {
			beforeWeight = rand.Float64()*(maxMaleWeight-minFemaleWeight) + minFemaleWeight
			height = rand.Float64()*(maxMaleHeight-maxFemaleHeight) + minFemaleHeight
		} else {
			beforeWeight = rand.Float64()*(maxMaleWeight-minMaleWeight) + minMaleWeight
			height = rand.Float64()*(maxMaleHeight-minMaleHeight) + minMaleHeight
		}

		if rand.Float64() < badMUSTThres {
			curWeight = maxFemaleWeight - (rand.Float64() * (maxMaleWeight - minFemaleWeight))
		} else {
			curWeight = maxMaleWeight - (rand.Float64() * (maxMaleWeight - minMaleWeight))
		}
	} else if gender == "female" {
		beforeWeight = rand.Float64()*(maxFemaleWeight-minFemaleWeight) + minFemaleWeight

		if rand.Float64() < badMUSTThres {
			curWeight = maxFemaleWeight - (rand.Float64() * (maxMaleWeight - minFemaleWeight))
			height = rand.Float64()*(maxMaleHeight-maxFemaleHeight) + minMaleHeight
		} else {
			curWeight = maxFemaleWeight - (rand.Float64() * (maxFemaleWeight - minFemaleWeight))
			height = rand.Float64()*(maxFemaleHeight-minFemaleHeight) + minFemaleHeight
		}
	} else {
		return nil, fmt.Errorf("`gender` has to be 'male' or 'female'!")
	}

	result["height"] = height
	result["weight_before"] = beforeWeight
	result["weight_current"] = curWeight
	q1Res, err := randomSample(choices, 1)
	if err != nil {
		return nil, err
	}
	q2Res, err := randomSample(choices, 1)
	if err != nil {
		return nil, err
	}
	result["must_q1"] = q1Res[0]
	result["must_q2"] = q2Res[0]

	return result, nil
}

// Fakes EQ5D5L data.  We'll go ahead and pick one random response
// for each domain in the survey - the final rating in the end
// is meant to be an integer, but can be a float if rand.Float64()
// returns a float lesser than eq5d5lThres.
func fakeEQ5D5LData() (map[string]any, error) {
	eq5d5lData, err := loadJson[map[string][]string](eq5d5lPath)
	if err != nil {
		return nil, err
	}

	eq5d5lSample := make(map[string]any)
	for k, v := range eq5d5lData {
		eq5d5lSample[k] = v[rand.IntN(len(v))]
	}
	rating := float64(rand.IntN(eq5d5lRatingMax))
	if rand.Float64() < eq5d5lThres {
		rating += rand.Float64()
	}
	eq5d5lSample["rating"] = rating

	return eq5d5lSample, nil
}
