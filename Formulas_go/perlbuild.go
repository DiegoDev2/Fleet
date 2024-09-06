package main

// PerlBuild - Perl builder
// Homepage: https://github.com/tokuhirom/Perl-Build

import (
	"fmt"
	
	"os/exec"
)

func installPerlBuild() {
	// Método 1: Descargar y extraer .tar.gz
	perlbuild_tar_url := "https://github.com/tokuhirom/Perl-Build/archive/refs/tags/1.34.tar.gz"
	perlbuild_cmd_tar := exec.Command("curl", "-L", perlbuild_tar_url, "-o", "package.tar.gz")
	err := perlbuild_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	perlbuild_zip_url := "https://github.com/tokuhirom/Perl-Build/archive/refs/tags/1.34.zip"
	perlbuild_cmd_zip := exec.Command("curl", "-L", perlbuild_zip_url, "-o", "package.zip")
	err = perlbuild_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	perlbuild_bin_url := "https://github.com/tokuhirom/Perl-Build/archive/refs/tags/1.34.bin"
	perlbuild_cmd_bin := exec.Command("curl", "-L", perlbuild_bin_url, "-o", "binary.bin")
	err = perlbuild_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	perlbuild_src_url := "https://github.com/tokuhirom/Perl-Build/archive/refs/tags/1.34.src.tar.gz"
	perlbuild_cmd_src := exec.Command("curl", "-L", perlbuild_src_url, "-o", "source.tar.gz")
	err = perlbuild_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	perlbuild_cmd_direct := exec.Command("./binary")
	err = perlbuild_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
