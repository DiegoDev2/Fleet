package main

// MecabJumandic - See mecab
// Homepage: https://taku910.github.io/mecab/

import (
	"fmt"
	
	"os/exec"
)

func installMecabJumandic() {
	// Método 1: Descargar y extraer .tar.gz
	mecabjumandic_tar_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/mecab/mecab-jumandic-7.0-20130310.tar.gz"
	mecabjumandic_cmd_tar := exec.Command("curl", "-L", mecabjumandic_tar_url, "-o", "package.tar.gz")
	err := mecabjumandic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mecabjumandic_zip_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/mecab/mecab-jumandic-7.0-20130310.zip"
	mecabjumandic_cmd_zip := exec.Command("curl", "-L", mecabjumandic_zip_url, "-o", "package.zip")
	err = mecabjumandic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mecabjumandic_bin_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/mecab/mecab-jumandic-7.0-20130310.bin"
	mecabjumandic_cmd_bin := exec.Command("curl", "-L", mecabjumandic_bin_url, "-o", "binary.bin")
	err = mecabjumandic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mecabjumandic_src_url := "https://www.mirrorservice.org/sites/distfiles.macports.org/mecab/mecab-jumandic-7.0-20130310.src.tar.gz"
	mecabjumandic_cmd_src := exec.Command("curl", "-L", mecabjumandic_src_url, "-o", "source.tar.gz")
	err = mecabjumandic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mecabjumandic_cmd_direct := exec.Command("./binary")
	err = mecabjumandic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: mecab")
exec.Command("latte", "install", "mecab")
}
