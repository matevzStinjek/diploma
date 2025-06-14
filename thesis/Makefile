.PHONY: all clean

# Main document name
MAIN = thesis

# Compiler
LATEX = pdflatex
BIBTEX = bibtex

# Default target
all: $(MAIN).pdf

# Compile the document
$(MAIN).pdf: $(MAIN).tex
	$(LATEX) $(MAIN)
	$(BIBTEX) $(MAIN)
	$(LATEX) $(MAIN)
	$(LATEX) $(MAIN)

# Clean temporary files
clean:
	rm -f *.aux *.log *.out *.toc *.lof *.lot *.bbl *.blg *.fls *.fdb_latexmk *.synctex.gz

# Clean everything except source files
cleanall: clean
	rm -f $(MAIN).pdf 