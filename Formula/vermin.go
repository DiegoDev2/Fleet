package main

// Vermin - Concurrently detect the minimum Python versions needed to run code
// Homepage: https://github.com/netromdk/vermin

import (
	"fmt"
	
	"os/exec"
)

func installVermin() {
	// Método 1: Descargar y extraer .tar.gz
	vermin_tar_url := "https://files.pythonhosted.org/packages/3d/26/7b871396c33006c445c25ef7da605ecbd6cef830d577b496d2b73a554f9d/vermin-1.6.0.tar.gz"
	vermin_cmd_tar := exec.Command("curl", "-L", vermin_tar_url, "-o", "package.tar.gz")
	err := vermin_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	vermin_zip_url := "https://files.pythonhosted.org/packages/3d/26/7b871396c33006c445c25ef7da605ecbd6cef830d577b496d2b73a554f9d/vermin-1.6.0.zip"
	vermin_cmd_zip := exec.Command("curl", "-L", vermin_zip_url, "-o", "package.zip")
	err = vermin_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	vermin_bin_url := "https://files.pythonhosted.org/packages/3d/26/7b871396c33006c445c25ef7da605ecbd6cef830d577b496d2b73a554f9d/vermin-1.6.0.bin"
	vermin_cmd_bin := exec.Command("curl", "-L", vermin_bin_url, "-o", "binary.bin")
	err = vermin_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	vermin_src_url := "https://files.pythonhosted.org/packages/3d/26/7b871396c33006c445c25ef7da605ecbd6cef830d577b496d2b73a554f9d/vermin-1.6.0.src.tar.gz"
	vermin_cmd_src := exec.Command("curl", "-L", vermin_src_url, "-o", "source.tar.gz")
	err = vermin_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	vermin_cmd_direct := exec.Command("./binary")
	err = vermin_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
