package main

// Mkvtomp4 - Convert mkv files to mp4
// Homepage: https://github.com/gavinbeatty/mkvtomp4/

import (
	"fmt"
	
	"os/exec"
)

func installMkvtomp4() {
	// Método 1: Descargar y extraer .tar.gz
	mkvtomp4_tar_url := "https://files.pythonhosted.org/packages/89/27/7367092f0d5530207e049afc76b167998dca2478a5c004018cf07e8a5653/mkvtomp4-2.0.tar.gz"
	mkvtomp4_cmd_tar := exec.Command("curl", "-L", mkvtomp4_tar_url, "-o", "package.tar.gz")
	err := mkvtomp4_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mkvtomp4_zip_url := "https://files.pythonhosted.org/packages/89/27/7367092f0d5530207e049afc76b167998dca2478a5c004018cf07e8a5653/mkvtomp4-2.0.zip"
	mkvtomp4_cmd_zip := exec.Command("curl", "-L", mkvtomp4_zip_url, "-o", "package.zip")
	err = mkvtomp4_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mkvtomp4_bin_url := "https://files.pythonhosted.org/packages/89/27/7367092f0d5530207e049afc76b167998dca2478a5c004018cf07e8a5653/mkvtomp4-2.0.bin"
	mkvtomp4_cmd_bin := exec.Command("curl", "-L", mkvtomp4_bin_url, "-o", "binary.bin")
	err = mkvtomp4_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mkvtomp4_src_url := "https://files.pythonhosted.org/packages/89/27/7367092f0d5530207e049afc76b167998dca2478a5c004018cf07e8a5653/mkvtomp4-2.0.src.tar.gz"
	mkvtomp4_cmd_src := exec.Command("curl", "-L", mkvtomp4_src_url, "-o", "source.tar.gz")
	err = mkvtomp4_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mkvtomp4_cmd_direct := exec.Command("./binary")
	err = mkvtomp4_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ffmpeg")
	exec.Command("latte", "install", "ffmpeg").Run()
	fmt.Println("Instalando dependencia: gpac")
	exec.Command("latte", "install", "gpac").Run()
	fmt.Println("Instalando dependencia: mkvtoolnix")
	exec.Command("latte", "install", "mkvtoolnix").Run()
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
