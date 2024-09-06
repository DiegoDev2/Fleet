package main

// FseventsTools - Command-line utilities for the FSEvents API
// Homepage: https://geoff.greer.fm/fsevents/

import (
	"fmt"
	
	"os/exec"
)

func installFseventsTools() {
	// Método 1: Descargar y extraer .tar.gz
	fseventstools_tar_url := "https://geoff.greer.fm/fsevents/releases/fsevents-tools-1.0.0.tar.gz"
	fseventstools_cmd_tar := exec.Command("curl", "-L", fseventstools_tar_url, "-o", "package.tar.gz")
	err := fseventstools_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	fseventstools_zip_url := "https://geoff.greer.fm/fsevents/releases/fsevents-tools-1.0.0.zip"
	fseventstools_cmd_zip := exec.Command("curl", "-L", fseventstools_zip_url, "-o", "package.zip")
	err = fseventstools_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	fseventstools_bin_url := "https://geoff.greer.fm/fsevents/releases/fsevents-tools-1.0.0.bin"
	fseventstools_cmd_bin := exec.Command("curl", "-L", fseventstools_bin_url, "-o", "binary.bin")
	err = fseventstools_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	fseventstools_src_url := "https://geoff.greer.fm/fsevents/releases/fsevents-tools-1.0.0.src.tar.gz"
	fseventstools_cmd_src := exec.Command("curl", "-L", fseventstools_src_url, "-o", "source.tar.gz")
	err = fseventstools_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	fseventstools_cmd_direct := exec.Command("./binary")
	err = fseventstools_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
}
