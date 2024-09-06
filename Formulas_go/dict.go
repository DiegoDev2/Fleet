package main

// Dict - Dictionary Server Protocol (RFC2229) client
// Homepage: https://dict.org/bin/Dict

import (
	"fmt"
	
	"os/exec"
)

func installDict() {
	// Método 1: Descargar y extraer .tar.gz
	dict_tar_url := "https://downloads.sourceforge.net/project/dict/dictd/dictd-1.13.1/dictd-1.13.1.tar.gz"
	dict_cmd_tar := exec.Command("curl", "-L", dict_tar_url, "-o", "package.tar.gz")
	err := dict_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dict_zip_url := "https://downloads.sourceforge.net/project/dict/dictd/dictd-1.13.1/dictd-1.13.1.zip"
	dict_cmd_zip := exec.Command("curl", "-L", dict_zip_url, "-o", "package.zip")
	err = dict_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dict_bin_url := "https://downloads.sourceforge.net/project/dict/dictd/dictd-1.13.1/dictd-1.13.1.bin"
	dict_cmd_bin := exec.Command("curl", "-L", dict_bin_url, "-o", "binary.bin")
	err = dict_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dict_src_url := "https://downloads.sourceforge.net/project/dict/dictd/dictd-1.13.1/dictd-1.13.1.src.tar.gz"
	dict_cmd_src := exec.Command("curl", "-L", dict_src_url, "-o", "source.tar.gz")
	err = dict_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dict_cmd_direct := exec.Command("./binary")
	err = dict_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtool")
exec.Command("latte", "install", "libtool")
	fmt.Println("Instalando dependencia: libmaa")
exec.Command("latte", "install", "libmaa")
}
