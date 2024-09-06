package main

// Libmaxminddb - C library for the MaxMind DB file format
// Homepage: https://github.com/maxmind/libmaxminddb

import (
	"fmt"
	
	"os/exec"
)

func installLibmaxminddb() {
	// Método 1: Descargar y extraer .tar.gz
	libmaxminddb_tar_url := "https://github.com/maxmind/libmaxminddb/releases/download/1.11.0/libmaxminddb-1.11.0.tar.gz"
	libmaxminddb_cmd_tar := exec.Command("curl", "-L", libmaxminddb_tar_url, "-o", "package.tar.gz")
	err := libmaxminddb_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libmaxminddb_zip_url := "https://github.com/maxmind/libmaxminddb/releases/download/1.11.0/libmaxminddb-1.11.0.zip"
	libmaxminddb_cmd_zip := exec.Command("curl", "-L", libmaxminddb_zip_url, "-o", "package.zip")
	err = libmaxminddb_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libmaxminddb_bin_url := "https://github.com/maxmind/libmaxminddb/releases/download/1.11.0/libmaxminddb-1.11.0.bin"
	libmaxminddb_cmd_bin := exec.Command("curl", "-L", libmaxminddb_bin_url, "-o", "binary.bin")
	err = libmaxminddb_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libmaxminddb_src_url := "https://github.com/maxmind/libmaxminddb/releases/download/1.11.0/libmaxminddb-1.11.0.src.tar.gz"
	libmaxminddb_cmd_src := exec.Command("curl", "-L", libmaxminddb_src_url, "-o", "source.tar.gz")
	err = libmaxminddb_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libmaxminddb_cmd_direct := exec.Command("./binary")
	err = libmaxminddb_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
	exec.Command("latte", "install", "autoconf").Run()
	fmt.Println("Instalando dependencia: automake")
	exec.Command("latte", "install", "automake").Run()
	fmt.Println("Instalando dependencia: libtool")
	exec.Command("latte", "install", "libtool").Run()
}
