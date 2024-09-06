package main

// Libeatmydata - LD_PRELOAD library and wrapper to transparently disable fsync and related calls
// Homepage: https://www.flamingspork.com/projects/libeatmydata/

import (
	"fmt"
	
	"os/exec"
)

func installLibeatmydata() {
	// Método 1: Descargar y extraer .tar.gz
	libeatmydata_tar_url := "https://github.com/stewartsmith/libeatmydata/releases/download/v131/libeatmydata-131.tar.gz"
	libeatmydata_cmd_tar := exec.Command("curl", "-L", libeatmydata_tar_url, "-o", "package.tar.gz")
	err := libeatmydata_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libeatmydata_zip_url := "https://github.com/stewartsmith/libeatmydata/releases/download/v131/libeatmydata-131.zip"
	libeatmydata_cmd_zip := exec.Command("curl", "-L", libeatmydata_zip_url, "-o", "package.zip")
	err = libeatmydata_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libeatmydata_bin_url := "https://github.com/stewartsmith/libeatmydata/releases/download/v131/libeatmydata-131.bin"
	libeatmydata_cmd_bin := exec.Command("curl", "-L", libeatmydata_bin_url, "-o", "binary.bin")
	err = libeatmydata_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libeatmydata_src_url := "https://github.com/stewartsmith/libeatmydata/releases/download/v131/libeatmydata-131.src.tar.gz"
	libeatmydata_cmd_src := exec.Command("curl", "-L", libeatmydata_src_url, "-o", "source.tar.gz")
	err = libeatmydata_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libeatmydata_cmd_direct := exec.Command("./binary")
	err = libeatmydata_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: autoconf-archive")
	exec.Command("latte", "install", "autoconf-archive").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
	fmt.Println("Instalando dependencia: coreutils")
	exec.Command("latte", "install", "coreutils").Run()
	fmt.Println("Instalando dependencia: strace")
	exec.Command("latte", "install", "strace").Run()
}
