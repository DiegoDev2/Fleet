package main

// GnustepBase - Library of general-purpose, non-graphical Objective C objects
// Homepage: https://github.com/gnustep/libs-base

import (
	"fmt"
	
	"os/exec"
)

func installGnustepBase() {
	// Método 1: Descargar y extraer .tar.gz
	gnustepbase_tar_url := "https://github.com/gnustep/libs-base/releases/download/base-1_30_0/gnustep-base-1.30.0.tar.gz"
	gnustepbase_cmd_tar := exec.Command("curl", "-L", gnustepbase_tar_url, "-o", "package.tar.gz")
	err := gnustepbase_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gnustepbase_zip_url := "https://github.com/gnustep/libs-base/releases/download/base-1_30_0/gnustep-base-1.30.0.zip"
	gnustepbase_cmd_zip := exec.Command("curl", "-L", gnustepbase_zip_url, "-o", "package.zip")
	err = gnustepbase_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gnustepbase_bin_url := "https://github.com/gnustep/libs-base/releases/download/base-1_30_0/gnustep-base-1.30.0.bin"
	gnustepbase_cmd_bin := exec.Command("curl", "-L", gnustepbase_bin_url, "-o", "binary.bin")
	err = gnustepbase_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gnustepbase_src_url := "https://github.com/gnustep/libs-base/releases/download/base-1_30_0/gnustep-base-1.30.0.src.tar.gz"
	gnustepbase_cmd_src := exec.Command("curl", "-L", gnustepbase_src_url, "-o", "source.tar.gz")
	err = gnustepbase_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gnustepbase_cmd_direct := exec.Command("./binary")
	err = gnustepbase_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnustep-make")
	exec.Command("latte", "install", "gnustep-make").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gmp")
	exec.Command("latte", "install", "gmp").Run()
	fmt.Println("Instalando dependencia: gnutls")
	exec.Command("latte", "install", "gnutls").Run()
	fmt.Println("Instalando dependencia: libobjc2")
	exec.Command("latte", "install", "libobjc2").Run()
	fmt.Println("Instalando dependencia: zstd")
	exec.Command("latte", "install", "zstd").Run()
}
