package main

// Jrtplib - Fully featured C++ Library for RTP (Real-time Transport Protocol)
// Homepage: https://research.edm.uhasselt.be/jori/jrtplib

import (
	"fmt"
	
	"os/exec"
)

func installJrtplib() {
	// Método 1: Descargar y extraer .tar.gz
	jrtplib_tar_url := "https://research.edm.uhasselt.be/jori/jrtplib/jrtplib-3.11.2.tar.bz2"
	jrtplib_cmd_tar := exec.Command("curl", "-L", jrtplib_tar_url, "-o", "package.tar.gz")
	err := jrtplib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	jrtplib_zip_url := "https://research.edm.uhasselt.be/jori/jrtplib/jrtplib-3.11.2.tar.bz2"
	jrtplib_cmd_zip := exec.Command("curl", "-L", jrtplib_zip_url, "-o", "package.zip")
	err = jrtplib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	jrtplib_bin_url := "https://research.edm.uhasselt.be/jori/jrtplib/jrtplib-3.11.2.tar.bz2"
	jrtplib_cmd_bin := exec.Command("curl", "-L", jrtplib_bin_url, "-o", "binary.bin")
	err = jrtplib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	jrtplib_src_url := "https://research.edm.uhasselt.be/jori/jrtplib/jrtplib-3.11.2.tar.bz2"
	jrtplib_cmd_src := exec.Command("curl", "-L", jrtplib_src_url, "-o", "source.tar.gz")
	err = jrtplib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	jrtplib_cmd_direct := exec.Command("./binary")
	err = jrtplib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: jthread")
exec.Command("latte", "install", "jthread")
}
