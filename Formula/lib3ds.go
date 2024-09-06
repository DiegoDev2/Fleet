package main

// Lib3ds - Library for managing 3D-Studio Release 3 and 4 '.3DS' files
// Homepage: https://code.google.com/archive/p/lib3ds/

import (
	"fmt"
	
	"os/exec"
)

func installLib3ds() {
	// Método 1: Descargar y extraer .tar.gz
	lib3ds_tar_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/lib3ds/lib3ds-1.3.0.zip"
	lib3ds_cmd_tar := exec.Command("curl", "-L", lib3ds_tar_url, "-o", "package.tar.gz")
	err := lib3ds_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	lib3ds_zip_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/lib3ds/lib3ds-1.3.0.zip"
	lib3ds_cmd_zip := exec.Command("curl", "-L", lib3ds_zip_url, "-o", "package.zip")
	err = lib3ds_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	lib3ds_bin_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/lib3ds/lib3ds-1.3.0.zip"
	lib3ds_cmd_bin := exec.Command("curl", "-L", lib3ds_bin_url, "-o", "binary.bin")
	err = lib3ds_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	lib3ds_src_url := "https://storage.googleapis.com/google-code-archive-downloads/v2/code.google.com/lib3ds/lib3ds-1.3.0.zip"
	lib3ds_cmd_src := exec.Command("curl", "-L", lib3ds_src_url, "-o", "source.tar.gz")
	err = lib3ds_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	lib3ds_cmd_direct := exec.Command("./binary")
	err = lib3ds_cmd_direct.Run()
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
