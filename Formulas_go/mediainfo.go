package main

// MediaInfo - Unified display of technical and tag data for audio/video
// Homepage: https://mediaarea.net/

import (
	"fmt"
	
	"os/exec"
)

func installMediaInfo() {
	// Método 1: Descargar y extraer .tar.gz
	mediainfo_tar_url := "https://mediaarea.net/download/binary/mediainfo/24.06/MediaInfo_CLI_24.06_GNU_FromSource.tar.bz2"
	mediainfo_cmd_tar := exec.Command("curl", "-L", mediainfo_tar_url, "-o", "package.tar.gz")
	err := mediainfo_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mediainfo_zip_url := "https://mediaarea.net/download/binary/mediainfo/24.06/MediaInfo_CLI_24.06_GNU_FromSource.tar.bz2"
	mediainfo_cmd_zip := exec.Command("curl", "-L", mediainfo_zip_url, "-o", "package.zip")
	err = mediainfo_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mediainfo_bin_url := "https://mediaarea.net/download/binary/mediainfo/24.06/MediaInfo_CLI_24.06_GNU_FromSource.tar.bz2"
	mediainfo_cmd_bin := exec.Command("curl", "-L", mediainfo_bin_url, "-o", "binary.bin")
	err = mediainfo_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mediainfo_src_url := "https://mediaarea.net/download/binary/mediainfo/24.06/MediaInfo_CLI_24.06_GNU_FromSource.tar.bz2"
	mediainfo_cmd_src := exec.Command("curl", "-L", mediainfo_src_url, "-o", "source.tar.gz")
	err = mediainfo_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mediainfo_cmd_direct := exec.Command("./binary")
	err = mediainfo_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: libmediainfo")
exec.Command("latte", "install", "libmediainfo")
	fmt.Println("Instalando dependencia: libzen")
exec.Command("latte", "install", "libzen")
}
