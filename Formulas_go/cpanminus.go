package main

// Cpanminus - Get, unpack, build, and install modules from CPAN
// Homepage: https://github.com/miyagawa/cpanminus

import (
	"fmt"
	
	"os/exec"
)

func installCpanminus() {
	// Método 1: Descargar y extraer .tar.gz
	cpanminus_tar_url := "https://cpan.metacpan.org/authors/id/M/MI/MIYAGAWA/App-cpanminus-1.7047.tar.gz"
	cpanminus_cmd_tar := exec.Command("curl", "-L", cpanminus_tar_url, "-o", "package.tar.gz")
	err := cpanminus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpanminus_zip_url := "https://cpan.metacpan.org/authors/id/M/MI/MIYAGAWA/App-cpanminus-1.7047.zip"
	cpanminus_cmd_zip := exec.Command("curl", "-L", cpanminus_zip_url, "-o", "package.zip")
	err = cpanminus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpanminus_bin_url := "https://cpan.metacpan.org/authors/id/M/MI/MIYAGAWA/App-cpanminus-1.7047.bin"
	cpanminus_cmd_bin := exec.Command("curl", "-L", cpanminus_bin_url, "-o", "binary.bin")
	err = cpanminus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpanminus_src_url := "https://cpan.metacpan.org/authors/id/M/MI/MIYAGAWA/App-cpanminus-1.7047.src.tar.gz"
	cpanminus_cmd_src := exec.Command("curl", "-L", cpanminus_src_url, "-o", "source.tar.gz")
	err = cpanminus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpanminus_cmd_direct := exec.Command("./binary")
	err = cpanminus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
