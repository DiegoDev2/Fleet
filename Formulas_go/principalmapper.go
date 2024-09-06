package main

// Principalmapper - Quickly evaluate IAM permissions in AWS
// Homepage: https://github.com/nccgroup/PMapper

import (
	"fmt"
	
	"os/exec"
)

func installPrincipalmapper() {
	// Método 1: Descargar y extraer .tar.gz
	principalmapper_tar_url := "https://files.pythonhosted.org/packages/3f/8c/3d2efe475e9244bd45e3a776ea8207f34a9bb15caaa02f6c95e473b2ada2/principalmapper-1.1.5.tar.gz"
	principalmapper_cmd_tar := exec.Command("curl", "-L", principalmapper_tar_url, "-o", "package.tar.gz")
	err := principalmapper_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	principalmapper_zip_url := "https://files.pythonhosted.org/packages/3f/8c/3d2efe475e9244bd45e3a776ea8207f34a9bb15caaa02f6c95e473b2ada2/principalmapper-1.1.5.zip"
	principalmapper_cmd_zip := exec.Command("curl", "-L", principalmapper_zip_url, "-o", "package.zip")
	err = principalmapper_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	principalmapper_bin_url := "https://files.pythonhosted.org/packages/3f/8c/3d2efe475e9244bd45e3a776ea8207f34a9bb15caaa02f6c95e473b2ada2/principalmapper-1.1.5.bin"
	principalmapper_cmd_bin := exec.Command("curl", "-L", principalmapper_bin_url, "-o", "binary.bin")
	err = principalmapper_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	principalmapper_src_url := "https://files.pythonhosted.org/packages/3f/8c/3d2efe475e9244bd45e3a776ea8207f34a9bb15caaa02f6c95e473b2ada2/principalmapper-1.1.5.src.tar.gz"
	principalmapper_cmd_src := exec.Command("curl", "-L", principalmapper_src_url, "-o", "source.tar.gz")
	err = principalmapper_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	principalmapper_cmd_direct := exec.Command("./binary")
	err = principalmapper_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
