package parser

type nucleotide byte // unexported type users can't construct their own.

type DNAStrand []nucleotide // because users can't construct their own
// nucleotides they also can't construct their
// own DNAStrands.

const (
	// These are exported values so they can use these nucleotides to construct a
	// DNAStrand with.
	A nucleotide = 'A'
	C nucleotide = 'C'
	G nucleotide = 'G'
	T nucleotide = 'T'
)

// This function allows them to actually construct a DNAstrand with a list of
//  nucleotides from the constants above.
func New(nts ...nucleotide) DNAStrand {
	return nts
}
