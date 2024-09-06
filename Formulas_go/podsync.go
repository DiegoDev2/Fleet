package main

// Podsync - Turn YouTube or Vimeo channels, users, or playlists into podcast feeds
// Homepage: https://github.com/mxpv/podsync

import (
	"fmt"
	
	"os/exec"
)

func installPodsync() {
	// Método 1: Descargar y extraer .tar.gz
	podsync_tar_url := "https://github.com/mxpv/podsync/archive/refs/tags/v2.7.0.tar.gz"
	podsync_cmd_tar := exec.Command("curl", "-L", podsync_tar_url, "-o", "package.tar.gz")
	err := podsync_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	podsync_zip_url := "https://github.com/mxpv/podsync/archive/refs/tags/v2.7.0.zip"
	podsync_cmd_zip := exec.Command("curl", "-L", podsync_zip_url, "-o", "package.zip")
	err = podsync_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	podsync_bin_url := "https://github.com/mxpv/podsync/archive/refs/tags/v2.7.0.bin"
	podsync_cmd_bin := exec.Command("curl", "-L", podsync_bin_url, "-o", "binary.bin")
	err = podsync_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	podsync_src_url := "https://github.com/mxpv/podsync/archive/refs/tags/v2.7.0.src.tar.gz"
	podsync_cmd_src := exec.Command("curl", "-L", podsync_src_url, "-o", "source.tar.gz")
	err = podsync_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	podsync_cmd_direct := exec.Command("./binary")
	err = podsync_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: go")
exec.Command("latte", "install", "go")
	fmt.Println("Instalando dependencia: ffmpeg")
exec.Command("latte", "install", "ffmpeg")
	fmt.Println("Instalando dependencia: yt-dlp")
exec.Command("latte", "install", "yt-dlp")
}
