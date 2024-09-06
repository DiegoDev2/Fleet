package main

// Libqalculate - Library for Qalculate! program
// Homepage: https://qalculate.github.io/

import (
	"fmt"
	
	"os/exec"
)

func installLibqalculate() {
	// Método 1: Descargar y extraer .tar.gz
	libqalculate_tar_url := "https://github.com/Qalculate/libqalculate/releases/download/v5.2.0/libqalculate-5.2.0.tar.gz"
	libqalculate_cmd_tar := exec.Command("curl", "-L", libqalculate_tar_url, "-o", "package.tar.gz")
	err := libqalculate_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libqalculate_zip_url := "https://github.com/Qalculate/libqalculate/releases/download/v5.2.0/libqalculate-5.2.0.zip"
	libqalculate_cmd_zip := exec.Command("curl", "-L", libqalculate_zip_url, "-o", "package.zip")
	err = libqalculate_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libqalculate_bin_url := "https://github.com/Qalculate/libqalculate/releases/download/v5.2.0/libqalculate-5.2.0.bin"
	libqalculate_cmd_bin := exec.Command("curl", "-L", libqalculate_bin_url, "-o", "binary.bin")
	err = libqalculate_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libqalculate_src_url := "https://github.com/Qalculate/libqalculate/releases/download/v5.2.0/libqalculate-5.2.0.src.tar.gz"
	libqalculate_cmd_src := exec.Command("curl", "-L", libqalculate_src_url, "-o", "source.tar.gz")
	err = libqalculate_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libqalculate_cmd_direct := exec.Command("./binary")
	err = libqalculate_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: intltool")
exec.Command("latte", "install", "intltool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: gettext")
exec.Command("latte", "install", "gettext")
	fmt.Println("Instalando dependencia: gnuplot")
exec.Command("latte", "install", "gnuplot")
	fmt.Println("Instalando dependencia: mpfr")
exec.Command("latte", "install", "mpfr")
	fmt.Println("Instalando dependencia: readline")
exec.Command("latte", "install", "readline")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
	fmt.Println("Instalando dependencia: perl-xml-parser")
exec.Command("latte", "install", "perl-xml-parser")
	fmt.Println("Instalando dependencia: gmp")
exec.Command("latte", "install", "gmp")
}
