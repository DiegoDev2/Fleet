package main

// OpenalSoft - Implementation of the OpenAL 3D audio API
// Homepage: https://openal-soft.org/

import (
	"fmt"
	
	"os/exec"
)

func installOpenalSoft() {
	// Método 1: Descargar y extraer .tar.gz
	openalsoft_tar_url := "https://openal-soft.org/openal-releases/openal-soft-1.23.1.tar.bz2"
	openalsoft_cmd_tar := exec.Command("curl", "-L", openalsoft_tar_url, "-o", "package.tar.gz")
	err := openalsoft_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	openalsoft_zip_url := "https://openal-soft.org/openal-releases/openal-soft-1.23.1.tar.bz2"
	openalsoft_cmd_zip := exec.Command("curl", "-L", openalsoft_zip_url, "-o", "package.zip")
	err = openalsoft_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	openalsoft_bin_url := "https://openal-soft.org/openal-releases/openal-soft-1.23.1.tar.bz2"
	openalsoft_cmd_bin := exec.Command("curl", "-L", openalsoft_bin_url, "-o", "binary.bin")
	err = openalsoft_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	openalsoft_src_url := "https://openal-soft.org/openal-releases/openal-soft-1.23.1.tar.bz2"
	openalsoft_cmd_src := exec.Command("curl", "-L", openalsoft_src_url, "-o", "source.tar.gz")
	err = openalsoft_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	openalsoft_cmd_direct := exec.Command("./binary")
	err = openalsoft_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
