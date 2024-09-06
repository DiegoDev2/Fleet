package main

// Airshare - Cross-platform content sharing in a local network
// Homepage: https://airshare.readthedocs.io/

import (
	"fmt"
	
	"os/exec"
)

func installAirshare() {
	// Método 1: Descargar y extraer .tar.gz
	airshare_tar_url := "https://files.pythonhosted.org/packages/cb/a2/d59c18cd6a143bf860c29acb70552b7351fd7e0f56213be86b624601106b/Airshare-0.1.6.tar.gz"
	airshare_cmd_tar := exec.Command("curl", "-L", airshare_tar_url, "-o", "package.tar.gz")
	err := airshare_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	airshare_zip_url := "https://files.pythonhosted.org/packages/cb/a2/d59c18cd6a143bf860c29acb70552b7351fd7e0f56213be86b624601106b/Airshare-0.1.6.zip"
	airshare_cmd_zip := exec.Command("curl", "-L", airshare_zip_url, "-o", "package.zip")
	err = airshare_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	airshare_bin_url := "https://files.pythonhosted.org/packages/cb/a2/d59c18cd6a143bf860c29acb70552b7351fd7e0f56213be86b624601106b/Airshare-0.1.6.bin"
	airshare_cmd_bin := exec.Command("curl", "-L", airshare_bin_url, "-o", "binary.bin")
	err = airshare_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	airshare_src_url := "https://files.pythonhosted.org/packages/cb/a2/d59c18cd6a143bf860c29acb70552b7351fd7e0f56213be86b624601106b/Airshare-0.1.6.src.tar.gz"
	airshare_cmd_src := exec.Command("curl", "-L", airshare_src_url, "-o", "source.tar.gz")
	err = airshare_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	airshare_cmd_direct := exec.Command("./binary")
	err = airshare_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
	exec.Command("latte", "install", "python@3.12").Run()
}
