package main

// Bioperl - Perl tools for bioinformatics, genomics and life science
// Homepage: https://bioperl.org

import (
	"fmt"
	
	"os/exec"
)

func installBioperl() {
	// Método 1: Descargar y extraer .tar.gz
	bioperl_tar_url := "https://cpan.metacpan.org/authors/id/C/CJ/CJFIELDS/BioPerl-1.7.8.tar.gz"
	bioperl_cmd_tar := exec.Command("curl", "-L", bioperl_tar_url, "-o", "package.tar.gz")
	err := bioperl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	bioperl_zip_url := "https://cpan.metacpan.org/authors/id/C/CJ/CJFIELDS/BioPerl-1.7.8.zip"
	bioperl_cmd_zip := exec.Command("curl", "-L", bioperl_zip_url, "-o", "package.zip")
	err = bioperl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	bioperl_bin_url := "https://cpan.metacpan.org/authors/id/C/CJ/CJFIELDS/BioPerl-1.7.8.bin"
	bioperl_cmd_bin := exec.Command("curl", "-L", bioperl_bin_url, "-o", "binary.bin")
	err = bioperl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	bioperl_src_url := "https://cpan.metacpan.org/authors/id/C/CJ/CJFIELDS/BioPerl-1.7.8.src.tar.gz"
	bioperl_cmd_src := exec.Command("curl", "-L", bioperl_src_url, "-o", "source.tar.gz")
	err = bioperl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	bioperl_cmd_direct := exec.Command("./binary")
	err = bioperl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cpanminus")
	exec.Command("latte", "install", "cpanminus").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: perl")
	exec.Command("latte", "install", "perl").Run()
}
