package main

// Cpm - Fast CPAN module installer
// Homepage: https://metacpan.org/pod/cpm

import (
	"fmt"
	
	"os/exec"
)

func installCpm() {
	// Método 1: Descargar y extraer .tar.gz
	cpm_tar_url := "https://cpan.metacpan.org/authors/id/S/SK/SKAJI/App-cpm-0.997017.tar.gz"
	cpm_cmd_tar := exec.Command("curl", "-L", cpm_tar_url, "-o", "package.tar.gz")
	err := cpm_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	cpm_zip_url := "https://cpan.metacpan.org/authors/id/S/SK/SKAJI/App-cpm-0.997017.zip"
	cpm_cmd_zip := exec.Command("curl", "-L", cpm_zip_url, "-o", "package.zip")
	err = cpm_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	cpm_bin_url := "https://cpan.metacpan.org/authors/id/S/SK/SKAJI/App-cpm-0.997017.bin"
	cpm_cmd_bin := exec.Command("curl", "-L", cpm_bin_url, "-o", "binary.bin")
	err = cpm_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	cpm_src_url := "https://cpan.metacpan.org/authors/id/S/SK/SKAJI/App-cpm-0.997017.src.tar.gz"
	cpm_cmd_src := exec.Command("curl", "-L", cpm_src_url, "-o", "source.tar.gz")
	err = cpm_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	cpm_cmd_direct := exec.Command("./binary")
	err = cpm_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: perl")
	exec.Command("latte", "install", "perl").Run()
}
