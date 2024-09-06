package main

// Hy - Dialect of Lisp that's embedded in Python
// Homepage: https://github.com/hylang/hy

import (
	"fmt"
	
	"os/exec"
)

func installHy() {
	// Método 1: Descargar y extraer .tar.gz
	hy_tar_url := "https://files.pythonhosted.org/packages/88/53/e92bfd8a36dc4a62e0922d409f703299eac8a0a74ed4db2106acad4f00a0/hy-0.29.0.tar.gz"
	hy_cmd_tar := exec.Command("curl", "-L", hy_tar_url, "-o", "package.tar.gz")
	err := hy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	hy_zip_url := "https://files.pythonhosted.org/packages/88/53/e92bfd8a36dc4a62e0922d409f703299eac8a0a74ed4db2106acad4f00a0/hy-0.29.0.zip"
	hy_cmd_zip := exec.Command("curl", "-L", hy_zip_url, "-o", "package.zip")
	err = hy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	hy_bin_url := "https://files.pythonhosted.org/packages/88/53/e92bfd8a36dc4a62e0922d409f703299eac8a0a74ed4db2106acad4f00a0/hy-0.29.0.bin"
	hy_cmd_bin := exec.Command("curl", "-L", hy_bin_url, "-o", "binary.bin")
	err = hy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	hy_src_url := "https://files.pythonhosted.org/packages/88/53/e92bfd8a36dc4a62e0922d409f703299eac8a0a74ed4db2106acad4f00a0/hy-0.29.0.src.tar.gz"
	hy_cmd_src := exec.Command("curl", "-L", hy_src_url, "-o", "source.tar.gz")
	err = hy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	hy_cmd_direct := exec.Command("./binary")
	err = hy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
