package main

// Mp3val - Program for MPEG audio stream validation
// Homepage: https://mp3val.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMp3val() {
	// Método 1: Descargar y extraer .tar.gz
	mp3val_tar_url := "https://downloads.sourceforge.net/project/mp3val/mp3val/mp3val%200.1.8/mp3val-0.1.8-src.tar.gz"
	mp3val_cmd_tar := exec.Command("curl", "-L", mp3val_tar_url, "-o", "package.tar.gz")
	err := mp3val_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mp3val_zip_url := "https://downloads.sourceforge.net/project/mp3val/mp3val/mp3val%200.1.8/mp3val-0.1.8-src.zip"
	mp3val_cmd_zip := exec.Command("curl", "-L", mp3val_zip_url, "-o", "package.zip")
	err = mp3val_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mp3val_bin_url := "https://downloads.sourceforge.net/project/mp3val/mp3val/mp3val%200.1.8/mp3val-0.1.8-src.bin"
	mp3val_cmd_bin := exec.Command("curl", "-L", mp3val_bin_url, "-o", "binary.bin")
	err = mp3val_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mp3val_src_url := "https://downloads.sourceforge.net/project/mp3val/mp3val/mp3val%200.1.8/mp3val-0.1.8-src.src.tar.gz"
	mp3val_cmd_src := exec.Command("curl", "-L", mp3val_src_url, "-o", "source.tar.gz")
	err = mp3val_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mp3val_cmd_direct := exec.Command("./binary")
	err = mp3val_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
