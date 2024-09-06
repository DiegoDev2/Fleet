package main

// Libtickit - Library for building interactive full-screen terminal programs
// Homepage: https://www.leonerd.org.uk/code/libtickit/

import (
	"fmt"
	
	"os/exec"
)

func installLibtickit() {
	// Método 1: Descargar y extraer .tar.gz
	libtickit_tar_url := "https://www.leonerd.org.uk/code/libtickit/libtickit-0.4.3.tar.gz"
	libtickit_cmd_tar := exec.Command("curl", "-L", libtickit_tar_url, "-o", "package.tar.gz")
	err := libtickit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libtickit_zip_url := "https://www.leonerd.org.uk/code/libtickit/libtickit-0.4.3.zip"
	libtickit_cmd_zip := exec.Command("curl", "-L", libtickit_zip_url, "-o", "package.zip")
	err = libtickit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libtickit_bin_url := "https://www.leonerd.org.uk/code/libtickit/libtickit-0.4.3.bin"
	libtickit_cmd_bin := exec.Command("curl", "-L", libtickit_bin_url, "-o", "binary.bin")
	err = libtickit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libtickit_src_url := "https://www.leonerd.org.uk/code/libtickit/libtickit-0.4.3.src.tar.gz"
	libtickit_cmd_src := exec.Command("curl", "-L", libtickit_src_url, "-o", "source.tar.gz")
	err = libtickit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libtickit_cmd_direct := exec.Command("./binary")
	err = libtickit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libtermkey")
exec.Command("latte", "install", "libtermkey")
	fmt.Println("Instalando dependencia: unibilium")
exec.Command("latte", "install", "unibilium")
}
