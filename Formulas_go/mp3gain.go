package main

// Mp3gain - Lossless mp3 normalizer with statistical analysis
// Homepage: https://mp3gain.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installMp3gain() {
	// Método 1: Descargar y extraer .tar.gz
	mp3gain_tar_url := "https://downloads.sourceforge.net/project/mp3gain/mp3gain/1.6.2/mp3gain-1_6_2-src.zip"
	mp3gain_cmd_tar := exec.Command("curl", "-L", mp3gain_tar_url, "-o", "package.tar.gz")
	err := mp3gain_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mp3gain_zip_url := "https://downloads.sourceforge.net/project/mp3gain/mp3gain/1.6.2/mp3gain-1_6_2-src.zip"
	mp3gain_cmd_zip := exec.Command("curl", "-L", mp3gain_zip_url, "-o", "package.zip")
	err = mp3gain_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mp3gain_bin_url := "https://downloads.sourceforge.net/project/mp3gain/mp3gain/1.6.2/mp3gain-1_6_2-src.zip"
	mp3gain_cmd_bin := exec.Command("curl", "-L", mp3gain_bin_url, "-o", "binary.bin")
	err = mp3gain_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mp3gain_src_url := "https://downloads.sourceforge.net/project/mp3gain/mp3gain/1.6.2/mp3gain-1_6_2-src.zip"
	mp3gain_cmd_src := exec.Command("curl", "-L", mp3gain_src_url, "-o", "source.tar.gz")
	err = mp3gain_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mp3gain_cmd_direct := exec.Command("./binary")
	err = mp3gain_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mpg123")
exec.Command("latte", "install", "mpg123")
}
