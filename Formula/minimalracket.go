package main

// MinimalRacket - Modern programming language in the Lisp/Scheme family
// Homepage: https://racket-lang.org/

import (
	"fmt"
	
	"os/exec"
)

func installMinimalRacket() {
	// Método 1: Descargar y extraer .tar.gz
	minimalracket_tar_url := "https://mirror.racket-lang.org/installers/8.13/racket-minimal-8.13-src.tgz"
	minimalracket_cmd_tar := exec.Command("curl", "-L", minimalracket_tar_url, "-o", "package.tar.gz")
	err := minimalracket_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	minimalracket_zip_url := "https://mirror.racket-lang.org/installers/8.13/racket-minimal-8.13-src.tgz"
	minimalracket_cmd_zip := exec.Command("curl", "-L", minimalracket_zip_url, "-o", "package.zip")
	err = minimalracket_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	minimalracket_bin_url := "https://mirror.racket-lang.org/installers/8.13/racket-minimal-8.13-src.tgz"
	minimalracket_cmd_bin := exec.Command("curl", "-L", minimalracket_bin_url, "-o", "binary.bin")
	err = minimalracket_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	minimalracket_src_url := "https://mirror.racket-lang.org/installers/8.13/racket-minimal-8.13-src.tgz"
	minimalracket_cmd_src := exec.Command("curl", "-L", minimalracket_src_url, "-o", "source.tar.gz")
	err = minimalracket_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	minimalracket_cmd_direct := exec.Command("./binary")
	err = minimalracket_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: openssl@1.1")
	exec.Command("latte", "install", "openssl@1.1").Run()
}
