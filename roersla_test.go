package roersla

import "testing"

func TestConjugateNoun(t *testing.T) {
    got := ConjugateNoun("hest", Nominative, Masculine)
    want := "hestur"
    if got != want {
        t.Errorf("got %q, want %q", got, want)
    }
}

func TestVerbForms(t *testing.T) {
    if ConjugateVerbPreterite("hava") != "havade" {
        t.Error("preteritum av 'hava' vart ikkje 'havade'")
    }
    if ConjugateVerbParticiple("kasta") != "kastad" {
        t.Error("partisipp av 'kasta' vart ikkje 'kastad'")
    }
}

func TestOrthography(t *testing.T) {
    if err := ValidateOrthography("abc ðef"); err == nil {
        t.Error("forventa feil på ulovleg 'ð'")
    }
    if err := ValidateOrthography("tak þak"); err == nil {
        t.Error("forventa feil på 'þ'")
    }
    if err := ValidateOrthography("tíd og tú"); err != nil {
        t.Errorf("gyldig tekst feila: %v", err)
    }
}
