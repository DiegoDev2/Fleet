package main

// PerlXmlParser - Perl module for parsing XML documents
// Homepage: https://github.com/cpan-authors/XML-Parser

import (
	"fmt"
	
	"os/exec"
)

func installPerlXmlParser() {
	// Método 1: Descargar y extraer .tar.gz
	perlxmlparser_tar_url := "https://cpan.metacpan.org/authors/id/T/TO/TODDR/XML-Parser-2.47.tar.gz"
	perlxmlparser_cmd_tar := exec.Command("curl", "-L", perlxmlparser_tar_url, "-o", "package.tar.gz")
	err := perlxmlparser_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	perlxmlparser_zip_url := "https://cpan.metacpan.org/authors/id/T/TO/TODDR/XML-Parser-2.47.zip"
	perlxmlparser_cmd_zip := exec.Command("curl", "-L", perlxmlparser_zip_url, "-o", "package.zip")
	err = perlxmlparser_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	perlxmlparser_bin_url := "https://cpan.metacpan.org/authors/id/T/TO/TODDR/XML-Parser-2.47.bin"
	perlxmlparser_cmd_bin := exec.Command("curl", "-L", perlxmlparser_bin_url, "-o", "binary.bin")
	err = perlxmlparser_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	perlxmlparser_src_url := "https://cpan.metacpan.org/authors/id/T/TO/TODDR/XML-Parser-2.47.src.tar.gz"
	perlxmlparser_cmd_src := exec.Command("curl", "-L", perlxmlparser_src_url, "-o", "source.tar.gz")
	err = perlxmlparser_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	perlxmlparser_cmd_direct := exec.Command("./binary")
	err = perlxmlparser_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
