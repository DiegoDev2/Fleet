package main

// RegexOpt - Perl-compatible regular expression optimizer
// Homepage: https://bisqwit.iki.fi/source/regexopt.html

import (
	"fmt"
	
	"os/exec"
)

func installRegexOpt() {
	// Método 1: Descargar y extraer .tar.gz
	regexopt_tar_url := "https://bisqwit.iki.fi/src/arch/regex-opt-1.2.4.tar.gz"
	regexopt_cmd_tar := exec.Command("curl", "-L", regexopt_tar_url, "-o", "package.tar.gz")
	err := regexopt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	regexopt_zip_url := "https://bisqwit.iki.fi/src/arch/regex-opt-1.2.4.zip"
	regexopt_cmd_zip := exec.Command("curl", "-L", regexopt_zip_url, "-o", "package.zip")
	err = regexopt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	regexopt_bin_url := "https://bisqwit.iki.fi/src/arch/regex-opt-1.2.4.bin"
	regexopt_cmd_bin := exec.Command("curl", "-L", regexopt_bin_url, "-o", "binary.bin")
	err = regexopt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	regexopt_src_url := "https://bisqwit.iki.fi/src/arch/regex-opt-1.2.4.src.tar.gz"
	regexopt_cmd_src := exec.Command("curl", "-L", regexopt_src_url, "-o", "source.tar.gz")
	err = regexopt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	regexopt_cmd_direct := exec.Command("./binary")
	err = regexopt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
