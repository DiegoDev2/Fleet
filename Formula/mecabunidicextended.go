package main

// MecabUnidicExtended - Extended morphological analyzer for MeCab
// Homepage: https://osdn.net/projects/unidic/

import (
	"fmt"
	
	"os/exec"
)

func installMecabUnidicExtended() {
	// Método 1: Descargar y extraer .tar.gz
	mecabunidicextended_tar_url := "https://dotsrc.dl.osdn.net/osdn/unidic/58338/unidic-mecab_kana-accent-2.1.2_src.zip"
	mecabunidicextended_cmd_tar := exec.Command("curl", "-L", mecabunidicextended_tar_url, "-o", "package.tar.gz")
	err := mecabunidicextended_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mecabunidicextended_zip_url := "https://dotsrc.dl.osdn.net/osdn/unidic/58338/unidic-mecab_kana-accent-2.1.2_src.zip"
	mecabunidicextended_cmd_zip := exec.Command("curl", "-L", mecabunidicextended_zip_url, "-o", "package.zip")
	err = mecabunidicextended_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mecabunidicextended_bin_url := "https://dotsrc.dl.osdn.net/osdn/unidic/58338/unidic-mecab_kana-accent-2.1.2_src.zip"
	mecabunidicextended_cmd_bin := exec.Command("curl", "-L", mecabunidicextended_bin_url, "-o", "binary.bin")
	err = mecabunidicextended_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mecabunidicextended_src_url := "https://dotsrc.dl.osdn.net/osdn/unidic/58338/unidic-mecab_kana-accent-2.1.2_src.zip"
	mecabunidicextended_cmd_src := exec.Command("curl", "-L", mecabunidicextended_src_url, "-o", "source.tar.gz")
	err = mecabunidicextended_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mecabunidicextended_cmd_direct := exec.Command("./binary")
	err = mecabunidicextended_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mecab")
	exec.Command("latte", "install", "mecab").Run()
}
