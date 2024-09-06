package main

// Carton - Perl module dependency manager (aka Bundler for Perl)
// Homepage: https://metacpan.org/pod/Carton

import (
	"fmt"
	
	"os/exec"
)

func installCarton() {
	// Método 1: Descargar y extraer .tar.gz
	carton_tar_url := "https://cpan.metacpan.org/authors/id/M/MI/MIYAGAWA/Carton-v1.0.35.tar.gz"
	carton_cmd_tar := exec.Command("curl", "-L", carton_tar_url, "-o", "package.tar.gz")
	err := carton_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	carton_zip_url := "https://cpan.metacpan.org/authors/id/M/MI/MIYAGAWA/Carton-v1.0.35.zip"
	carton_cmd_zip := exec.Command("curl", "-L", carton_zip_url, "-o", "package.zip")
	err = carton_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	carton_bin_url := "https://cpan.metacpan.org/authors/id/M/MI/MIYAGAWA/Carton-v1.0.35.bin"
	carton_cmd_bin := exec.Command("curl", "-L", carton_bin_url, "-o", "binary.bin")
	err = carton_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	carton_src_url := "https://cpan.metacpan.org/authors/id/M/MI/MIYAGAWA/Carton-v1.0.35.src.tar.gz"
	carton_cmd_src := exec.Command("curl", "-L", carton_src_url, "-o", "source.tar.gz")
	err = carton_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	carton_cmd_direct := exec.Command("./binary")
	err = carton_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: perl")
exec.Command("latte", "install", "perl")
}
