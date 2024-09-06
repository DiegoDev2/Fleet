package main

// MecabKoDic - See mecab
// Homepage: https://bitbucket.org/eunjeon/mecab-ko-dic

import (
	"fmt"
	
	"os/exec"
)

func installMecabKoDic() {
	// Método 1: Descargar y extraer .tar.gz
	mecabkodic_tar_url := "https://bitbucket.org/eunjeon/mecab-ko-dic/downloads/mecab-ko-dic-2.1.1-20180720.tar.gz"
	mecabkodic_cmd_tar := exec.Command("curl", "-L", mecabkodic_tar_url, "-o", "package.tar.gz")
	err := mecabkodic_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	mecabkodic_zip_url := "https://bitbucket.org/eunjeon/mecab-ko-dic/downloads/mecab-ko-dic-2.1.1-20180720.zip"
	mecabkodic_cmd_zip := exec.Command("curl", "-L", mecabkodic_zip_url, "-o", "package.zip")
	err = mecabkodic_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	mecabkodic_bin_url := "https://bitbucket.org/eunjeon/mecab-ko-dic/downloads/mecab-ko-dic-2.1.1-20180720.bin"
	mecabkodic_cmd_bin := exec.Command("curl", "-L", mecabkodic_bin_url, "-o", "binary.bin")
	err = mecabkodic_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	mecabkodic_src_url := "https://bitbucket.org/eunjeon/mecab-ko-dic/downloads/mecab-ko-dic-2.1.1-20180720.src.tar.gz"
	mecabkodic_cmd_src := exec.Command("curl", "-L", mecabkodic_src_url, "-o", "source.tar.gz")
	err = mecabkodic_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	mecabkodic_cmd_direct := exec.Command("./binary")
	err = mecabkodic_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: autoconf")
exec.Command("latte", "install", "autoconf")
	fmt.Println("Instalando dependencia: automake")
exec.Command("latte", "install", "automake")
	fmt.Println("Instalando dependencia: mecab-ko")
exec.Command("latte", "install", "mecab-ko")
}
