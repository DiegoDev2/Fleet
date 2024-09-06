package main

// Btparse - BibTeX utility libraries
// Homepage: https://metacpan.org/pod/distribution/Text-BibTeX/btparse/doc/btparse.pod

import (
	"fmt"
	
	"os/exec"
)

func installBtparse() {
	// Método 1: Descargar y extraer .tar.gz
	btparse_tar_url := "https://cpan.metacpan.org/authors/id/A/AM/AMBS/btparse/btparse-0.35.tar.gz"
	btparse_cmd_tar := exec.Command("curl", "-L", btparse_tar_url, "-o", "package.tar.gz")
	err := btparse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	btparse_zip_url := "https://cpan.metacpan.org/authors/id/A/AM/AMBS/btparse/btparse-0.35.zip"
	btparse_cmd_zip := exec.Command("curl", "-L", btparse_zip_url, "-o", "package.zip")
	err = btparse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	btparse_bin_url := "https://cpan.metacpan.org/authors/id/A/AM/AMBS/btparse/btparse-0.35.bin"
	btparse_cmd_bin := exec.Command("curl", "-L", btparse_bin_url, "-o", "binary.bin")
	err = btparse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	btparse_src_url := "https://cpan.metacpan.org/authors/id/A/AM/AMBS/btparse/btparse-0.35.src.tar.gz"
	btparse_cmd_src := exec.Command("curl", "-L", btparse_src_url, "-o", "source.tar.gz")
	err = btparse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	btparse_cmd_direct := exec.Command("./binary")
	err = btparse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
