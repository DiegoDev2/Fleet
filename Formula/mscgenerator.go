package main

// MscGenerator - Draws signalling charts from textual description
// Homepage: https://gitlab.com/msc-generator/msc-generator

import (
	"fmt"
	
	"os/exec"
)

func installMscGenerator() {
	// Método 1: Descargar y extraer .tar.gz
	mscgenerator_tar_url := "https://gitlab.com/api/v4/projects/31167732/packages/generic/msc-generator/8.6.2/msc-generator-8.6.2.tar.gz"
	mscgenerator_cmd_tar := exec.Command("curl", "-L", mscgenerator_tar_url, "-o", "package.tar.gz")
	err := mscgenerator_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mscgenerator_zip_url := "https://gitlab.com/api/v4/projects/31167732/packages/generic/msc-generator/8.6.2/msc-generator-8.6.2.zip"
	mscgenerator_cmd_zip := exec.Command("curl", "-L", mscgenerator_zip_url, "-o", "package.zip")
	err = mscgenerator_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mscgenerator_bin_url := "https://gitlab.com/api/v4/projects/31167732/packages/generic/msc-generator/8.6.2/msc-generator-8.6.2.bin"
	mscgenerator_cmd_bin := exec.Command("curl", "-L", mscgenerator_bin_url, "-o", "binary.bin")
	err = mscgenerator_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mscgenerator_src_url := "https://gitlab.com/api/v4/projects/31167732/packages/generic/msc-generator/8.6.2/msc-generator-8.6.2.src.tar.gz"
	mscgenerator_cmd_src := exec.Command("curl", "-L", mscgenerator_src_url, "-o", "source.tar.gz")
	err = mscgenerator_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mscgenerator_cmd_direct := exec.Command("./binary")
	err = mscgenerator_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: bison")
	exec.Command("latte", "install", "bison").Run()
	fmt.Println("Instalando dependencia: help2man")
	exec.Command("latte", "install", "help2man").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: cairo")
	exec.Command("latte", "install", "cairo").Run()
	fmt.Println("Instalando dependencia: fontconfig")
	exec.Command("latte", "install", "fontconfig").Run()
	fmt.Println("Instalando dependencia: gcc")
	exec.Command("latte", "install", "gcc").Run()
	fmt.Println("Instalando dependencia: glpk")
	exec.Command("latte", "install", "glpk").Run()
	fmt.Println("Instalando dependencia: graphviz")
	exec.Command("latte", "install", "graphviz").Run()
	fmt.Println("Instalando dependencia: libpng")
	exec.Command("latte", "install", "libpng").Run()
	fmt.Println("Instalando dependencia: sdl2")
	exec.Command("latte", "install", "sdl2").Run()
	fmt.Println("Instalando dependencia: tinyxml2")
	exec.Command("latte", "install", "tinyxml2").Run()
	fmt.Println("Instalando dependencia: gnu-sed")
	exec.Command("latte", "install", "gnu-sed").Run()
	fmt.Println("Instalando dependencia: make")
	exec.Command("latte", "install", "make").Run()
	fmt.Println("Instalando dependencia: mesa")
	exec.Command("latte", "install", "mesa").Run()
}
