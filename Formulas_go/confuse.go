package main

// Confuse - Configuration file parser library written in C
// Homepage: https://github.com/libconfuse/libconfuse

import (
	"fmt"
	
	"os/exec"
)

func installConfuse() {
	// Método 1: Descargar y extraer .tar.gz
	confuse_tar_url := "https://github.com/libconfuse/libconfuse/releases/download/v3.3/confuse-3.3.tar.xz"
	confuse_cmd_tar := exec.Command("curl", "-L", confuse_tar_url, "-o", "package.tar.gz")
	err := confuse_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	confuse_zip_url := "https://github.com/libconfuse/libconfuse/releases/download/v3.3/confuse-3.3.tar.xz"
	confuse_cmd_zip := exec.Command("curl", "-L", confuse_zip_url, "-o", "package.zip")
	err = confuse_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	confuse_bin_url := "https://github.com/libconfuse/libconfuse/releases/download/v3.3/confuse-3.3.tar.xz"
	confuse_cmd_bin := exec.Command("curl", "-L", confuse_bin_url, "-o", "binary.bin")
	err = confuse_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	confuse_src_url := "https://github.com/libconfuse/libconfuse/releases/download/v3.3/confuse-3.3.tar.xz"
	confuse_cmd_src := exec.Command("curl", "-L", confuse_src_url, "-o", "source.tar.gz")
	err = confuse_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	confuse_cmd_direct := exec.Command("./binary")
	err = confuse_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
