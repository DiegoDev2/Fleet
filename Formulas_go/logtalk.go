package main

// Logtalk - Declarative object-oriented logic programming language
// Homepage: https://logtalk.org/

import (
	"fmt"
	
	"os/exec"
)

func installLogtalk() {
	// Método 1: Descargar y extraer .tar.gz
	logtalk_tar_url := "https://github.com/LogtalkDotOrg/logtalk3/archive/refs/tags/lgt3820stable.tar.gz"
	logtalk_cmd_tar := exec.Command("curl", "-L", logtalk_tar_url, "-o", "package.tar.gz")
	err := logtalk_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	logtalk_zip_url := "https://github.com/LogtalkDotOrg/logtalk3/archive/refs/tags/lgt3820stable.zip"
	logtalk_cmd_zip := exec.Command("curl", "-L", logtalk_zip_url, "-o", "package.zip")
	err = logtalk_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	logtalk_bin_url := "https://github.com/LogtalkDotOrg/logtalk3/archive/refs/tags/lgt3820stable.bin"
	logtalk_cmd_bin := exec.Command("curl", "-L", logtalk_bin_url, "-o", "binary.bin")
	err = logtalk_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	logtalk_src_url := "https://github.com/LogtalkDotOrg/logtalk3/archive/refs/tags/lgt3820stable.src.tar.gz"
	logtalk_cmd_src := exec.Command("curl", "-L", logtalk_src_url, "-o", "source.tar.gz")
	err = logtalk_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	logtalk_cmd_direct := exec.Command("./binary")
	err = logtalk_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: gnu-prolog")
exec.Command("latte", "install", "gnu-prolog")
}
