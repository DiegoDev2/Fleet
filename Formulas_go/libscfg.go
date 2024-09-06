package main

// Libscfg - C library for scfg
// Homepage: https://git.sr.ht/~emersion/libscfg

import (
	"fmt"
	
	"os/exec"
)

func installLibscfg() {
	// Método 1: Descargar y extraer .tar.gz
	libscfg_tar_url := "https://git.sr.ht/~emersion/libscfg/archive/v0.1.1.tar.gz"
	libscfg_cmd_tar := exec.Command("curl", "-L", libscfg_tar_url, "-o", "package.tar.gz")
	err := libscfg_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	libscfg_zip_url := "https://git.sr.ht/~emersion/libscfg/archive/v0.1.1.zip"
	libscfg_cmd_zip := exec.Command("curl", "-L", libscfg_zip_url, "-o", "package.zip")
	err = libscfg_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	libscfg_bin_url := "https://git.sr.ht/~emersion/libscfg/archive/v0.1.1.bin"
	libscfg_cmd_bin := exec.Command("curl", "-L", libscfg_bin_url, "-o", "binary.bin")
	err = libscfg_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	libscfg_src_url := "https://git.sr.ht/~emersion/libscfg/archive/v0.1.1.src.tar.gz"
	libscfg_cmd_src := exec.Command("curl", "-L", libscfg_src_url, "-o", "source.tar.gz")
	err = libscfg_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	libscfg_cmd_direct := exec.Command("./binary")
	err = libscfg_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: meson")
exec.Command("latte", "install", "meson")
	fmt.Println("Instalando dependencia: ninja")
exec.Command("latte", "install", "ninja")
}
