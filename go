#!/bin/tcsh

dvips -Ppdf -G0 Newton
ps2pdf -dPDFSETTINGS=/printer -dCompatibilityLevel=1.3 -dMaxSubsetPct=100 -dSubsetFonts=true -dEmbedAllFonts=true Newton.ps
