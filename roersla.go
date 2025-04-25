package roersla

import (
    "encoding/csv"
    "fmt"
    "io"
    "os"
    "regexp"
    "strings"
)

// Kasus
type Case string

const (
    Nominative  Case = "nominativ"
    Accusative  Case = "akkusativ"
    Dative      Case = "dativ"
    Genitive    Case = "genitiv"
)

// Kjønn
type Gender string

const (
    Masculine Gender = "hannkyn"
    Feminine  Gender = "hokyn"
    Neuter    Gender = "inkjekyn"
)

var nounEndings = map[Case]map[Gender]string{
    Nominative: {
        Masculine: "ur",
        Feminine:  "a",
        Neuter:    "",
    },
    Accusative: {
        Masculine: "in",
        Feminine:  "a",
        Neuter:    "i",
    },
    Dative: {
        Masculine: "um",
        Feminine:  "um",
        Neuter:    "um",
    },
    Genitive: {
        Masculine: "s",
        Feminine:  "ar",
        Neuter:    "s",
    },
}

// ConjugateNoun bøyer eit substantiv etter stam, kasus og kjønn.
func ConjugateNoun(stem string, c Case, g Gender) string {
    if endings, ok := nounEndings[c]; ok {
        if e, ok2 := endings[g]; ok2 {
            return stem + e
        }
    }
    return stem
}

// ConjugateVerbPreterite legg på -de
func ConjugateVerbPreterite(stem string) string {
    return stem + "de"
}

// ConjugateVerbParticiple legg på -d
func ConjugateVerbParticiple(stem string) string {
    return stem + "d"
}

// Ortografi­erstatningar
func ReplaceOrthography(text string) string {
    r := strings.NewReplacer("ð", "d", "þ", "t")
    return r.Replace(text)
}

var (
    illegalDRE = regexp.MustCompile(`\b[^\s]ð[^\s]*\b(?<!Óðinn|Harðang|Niðarós)`)
    illegalTRE = regexp.MustCompile(`þ`)
)

// ValidateOrthography feilar om ulovlege teikn finn
func ValidateOrthography(text string) error {
    if illegalDRE.MatchString(text) {
        return fmt.Errorf("illegal 'ð' utanom tillat namn")
    }
    if illegalTRE.MatchString(text) {
        return fmt.Errorf("illegal 'þ' – bruk 't'")
    }
    return nil
}

// LoadMiniDictionary les inn CSV med Faroese,Djúpnorsk,Bokmål
func LoadMiniDictionary(path string) (map[string]string, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    r := csv.NewReader(f)
    dict := make(map[string]string)
    for {
        rec, err := r.Read()
        if err == io.EOF {
            break
        }
        if err != nil {
            return nil, err
        }
        if len(rec) < 2 {
            continue
        }
        far := strings.TrimSpace(rec[0])
        djp := strings.TrimSpace(rec[1])
        dict[far] = djp
    }
    return dict, nil
}
