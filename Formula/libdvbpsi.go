package main

// Libdvbpsi - Library to decode/generate MPEG TS and DVB PSI tables
// Homepage: https://www.videolan.org/developers/libdvbpsi.html

import (
	"fmt"
	
	"os/exec"
)

func installLibdvbpsi() {
	// Método 1: Descargar y extraer .tar.gz
	libdvbpsi_tar_url := "https://get.videolan.org/libdvbpsi/1.3.3/libdvbpsi-1.3.3.tar.bz2"
	libdvbpsi_cmd_tar := exec.Command("curl", "-L", libdvbpsi_tar_url, "-o", "package.tar.gz")
	err := libdvbpsi_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libdvbpsi_zip_url := "https://get.videolan.org/libdvbpsi/1.3.3/libdvbpsi-1.3.3.tar.bz2"
	libdvbpsi_cmd_zip := exec.Command("curl", "-L", libdvbpsi_zip_url, "-o", "package.zip")
	err = libdvbpsi_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libdvbpsi_bin_url := "https://get.videolan.org/libdvbpsi/1.3.3/libdvbpsi-1.3.3.tar.bz2"
	libdvbpsi_cmd_bin := exec.Command("curl", "-L", libdvbpsi_bin_url, "-o", "binary.bin")
	err = libdvbpsi_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libdvbpsi_src_url := "https://get.videolan.org/libdvbpsi/1.3.3/libdvbpsi-1.3.3.tar.bz2"
	libdvbpsi_cmd_src := exec.Command("curl", "-L", libdvbpsi_src_url, "-o", "source.tar.gz")
	err = libdvbpsi_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libdvbpsi_cmd_direct := exec.Command("./binary")
	err = libdvbpsi_cmd_direct.Run()
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
