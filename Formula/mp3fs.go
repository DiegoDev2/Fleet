package main

// Mp3fs - Read-only FUSE file system: transcodes audio formats to MP3
// Homepage: https://khenriks.github.io/mp3fs/

import (
	"fmt"
	
	"os/exec"
)

func installMp3fs() {
	// Método 1: Descargar y extraer .tar.gz
	mp3fs_tar_url := "https://github.com/khenriks/mp3fs/releases/download/v1.1.1/mp3fs-1.1.1.tar.gz"
	mp3fs_cmd_tar := exec.Command("curl", "-L", mp3fs_tar_url, "-o", "package.tar.gz")
	err := mp3fs_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mp3fs_zip_url := "https://github.com/khenriks/mp3fs/releases/download/v1.1.1/mp3fs-1.1.1.zip"
	mp3fs_cmd_zip := exec.Command("curl", "-L", mp3fs_zip_url, "-o", "package.zip")
	err = mp3fs_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mp3fs_bin_url := "https://github.com/khenriks/mp3fs/releases/download/v1.1.1/mp3fs-1.1.1.bin"
	mp3fs_cmd_bin := exec.Command("curl", "-L", mp3fs_bin_url, "-o", "binary.bin")
	err = mp3fs_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mp3fs_src_url := "https://github.com/khenriks/mp3fs/releases/download/v1.1.1/mp3fs-1.1.1.src.tar.gz"
	mp3fs_cmd_src := exec.Command("curl", "-L", mp3fs_src_url, "-o", "source.tar.gz")
	err = mp3fs_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mp3fs_cmd_direct := exec.Command("./binary")
	err = mp3fs_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: flac")
	exec.Command("latte", "install", "flac").Run()
	fmt.Println("Instalando dependencia: lame")
	exec.Command("latte", "install", "lame").Run()
	fmt.Println("Instalando dependencia: libfuse@2")
	exec.Command("latte", "install", "libfuse@2").Run()
	fmt.Println("Instalando dependencia: libid3tag")
	exec.Command("latte", "install", "libid3tag").Run()
	fmt.Println("Instalando dependencia: libvorbis")
	exec.Command("latte", "install", "libvorbis").Run()
}
