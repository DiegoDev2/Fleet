package main

// Mp3splt - Command-line interface to split MP3 and Ogg Vorbis files
// Homepage: https://mp3splt.sourceforge.net/mp3splt_page/home.php

import (
	"fmt"
	
	"os/exec"
)

func installMp3splt() {
	// Método 1: Descargar y extraer .tar.gz
	mp3splt_tar_url := "https://downloads.sourceforge.net/project/mp3splt/mp3splt/2.6.2/mp3splt-2.6.2.tar.gz"
	mp3splt_cmd_tar := exec.Command("curl", "-L", mp3splt_tar_url, "-o", "package.tar.gz")
	err := mp3splt_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mp3splt_zip_url := "https://downloads.sourceforge.net/project/mp3splt/mp3splt/2.6.2/mp3splt-2.6.2.zip"
	mp3splt_cmd_zip := exec.Command("curl", "-L", mp3splt_zip_url, "-o", "package.zip")
	err = mp3splt_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mp3splt_bin_url := "https://downloads.sourceforge.net/project/mp3splt/mp3splt/2.6.2/mp3splt-2.6.2.bin"
	mp3splt_cmd_bin := exec.Command("curl", "-L", mp3splt_bin_url, "-o", "binary.bin")
	err = mp3splt_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mp3splt_src_url := "https://downloads.sourceforge.net/project/mp3splt/mp3splt/2.6.2/mp3splt-2.6.2.src.tar.gz"
	mp3splt_cmd_src := exec.Command("curl", "-L", mp3splt_src_url, "-o", "source.tar.gz")
	err = mp3splt_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mp3splt_cmd_direct := exec.Command("./binary")
	err = mp3splt_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libmp3splt")
	exec.Command("latte", "install", "libmp3splt").Run()
}
