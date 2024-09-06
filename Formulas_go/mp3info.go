package main

// Mp3info - MP3 technical info viewer and ID3 1.x tag editor
// Homepage: https://www.ibiblio.org/mp3info/

import (
	"fmt"
	
	"os/exec"
)

func installMp3info() {
	// Método 1: Descargar y extraer .tar.gz
	mp3info_tar_url := "https://www.ibiblio.org/pub/linux/apps/sound/mp3-utils/mp3info/mp3info-0.8.5a.tgz"
	mp3info_cmd_tar := exec.Command("curl", "-L", mp3info_tar_url, "-o", "package.tar.gz")
	err := mp3info_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mp3info_zip_url := "https://www.ibiblio.org/pub/linux/apps/sound/mp3-utils/mp3info/mp3info-0.8.5a.tgz"
	mp3info_cmd_zip := exec.Command("curl", "-L", mp3info_zip_url, "-o", "package.zip")
	err = mp3info_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mp3info_bin_url := "https://www.ibiblio.org/pub/linux/apps/sound/mp3-utils/mp3info/mp3info-0.8.5a.tgz"
	mp3info_cmd_bin := exec.Command("curl", "-L", mp3info_bin_url, "-o", "binary.bin")
	err = mp3info_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mp3info_src_url := "https://www.ibiblio.org/pub/linux/apps/sound/mp3-utils/mp3info/mp3info-0.8.5a.tgz"
	mp3info_cmd_src := exec.Command("curl", "-L", mp3info_src_url, "-o", "source.tar.gz")
	err = mp3info_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mp3info_cmd_direct := exec.Command("./binary")
	err = mp3info_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
