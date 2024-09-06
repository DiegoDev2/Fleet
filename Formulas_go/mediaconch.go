package main

// Mediaconch - Conformance checker and technical metadata reporter
// Homepage: https://mediaarea.net/MediaConch

import (
	"fmt"
	
	"os/exec"
)

func installMediaconch() {
	// Método 1: Descargar y extraer .tar.gz
	mediaconch_tar_url := "https://mediaarea.net/download/binary/mediaconch/24.06/MediaConch_CLI_24.06_GNU_FromSource.tar.bz2"
	mediaconch_cmd_tar := exec.Command("curl", "-L", mediaconch_tar_url, "-o", "package.tar.gz")
	err := mediaconch_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mediaconch_zip_url := "https://mediaarea.net/download/binary/mediaconch/24.06/MediaConch_CLI_24.06_GNU_FromSource.tar.bz2"
	mediaconch_cmd_zip := exec.Command("curl", "-L", mediaconch_zip_url, "-o", "package.zip")
	err = mediaconch_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mediaconch_bin_url := "https://mediaarea.net/download/binary/mediaconch/24.06/MediaConch_CLI_24.06_GNU_FromSource.tar.bz2"
	mediaconch_cmd_bin := exec.Command("curl", "-L", mediaconch_bin_url, "-o", "binary.bin")
	err = mediaconch_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mediaconch_src_url := "https://mediaarea.net/download/binary/mediaconch/24.06/MediaConch_CLI_24.06_GNU_FromSource.tar.bz2"
	mediaconch_cmd_src := exec.Command("curl", "-L", mediaconch_src_url, "-o", "source.tar.gz")
	err = mediaconch_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mediaconch_cmd_direct := exec.Command("./binary")
	err = mediaconch_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: jansson")
exec.Command("latte", "install", "jansson")
	fmt.Println("Instalando dependencia: libevent")
exec.Command("latte", "install", "libevent")
	fmt.Println("Instalando dependencia: sqlite")
exec.Command("latte", "install", "sqlite")
}
