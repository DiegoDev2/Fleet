package main

// MecabUnidic - Morphological analyzer for MeCab
// Homepage: https://osdn.net/projects/unidic/

import (
	"fmt"
	
	"os/exec"
)

func installMecabUnidic() {
	// Método 1: Descargar y extraer .tar.gz
	mecabunidic_tar_url := "https://dotsrc.dl.osdn.net/osdn/unidic/58338/unidic-mecab-2.1.2_src.zip"
	mecabunidic_cmd_tar := exec.Command("curl", "-L", mecabunidic_tar_url, "-o", "package.tar.gz")
	err := mecabunidic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mecabunidic_zip_url := "https://dotsrc.dl.osdn.net/osdn/unidic/58338/unidic-mecab-2.1.2_src.zip"
	mecabunidic_cmd_zip := exec.Command("curl", "-L", mecabunidic_zip_url, "-o", "package.zip")
	err = mecabunidic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mecabunidic_bin_url := "https://dotsrc.dl.osdn.net/osdn/unidic/58338/unidic-mecab-2.1.2_src.zip"
	mecabunidic_cmd_bin := exec.Command("curl", "-L", mecabunidic_bin_url, "-o", "binary.bin")
	err = mecabunidic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mecabunidic_src_url := "https://dotsrc.dl.osdn.net/osdn/unidic/58338/unidic-mecab-2.1.2_src.zip"
	mecabunidic_cmd_src := exec.Command("curl", "-L", mecabunidic_src_url, "-o", "source.tar.gz")
	err = mecabunidic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mecabunidic_cmd_direct := exec.Command("./binary")
	err = mecabunidic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mecab")
	exec.Command("latte", "install", "mecab").Run()
}
