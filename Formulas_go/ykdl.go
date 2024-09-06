package main

// Ykdl - Video downloader that focus on China mainland video sites
// Homepage: https://github.com/SeaHOH/ykdl

import (
	"fmt"
	
	"os/exec"
)

func installYkdl() {
	// Método 1: Descargar y extraer .tar.gz
	ykdl_tar_url := "https://files.pythonhosted.org/packages/f2/27/f4e7616a139c84a04edb7778db2b3cfb77348ab73020ff232b6551fa8bdd/ykdl-1.8.2.tar.gz"
	ykdl_cmd_tar := exec.Command("curl", "-L", ykdl_tar_url, "-o", "package.tar.gz")
	err := ykdl_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	ykdl_zip_url := "https://files.pythonhosted.org/packages/f2/27/f4e7616a139c84a04edb7778db2b3cfb77348ab73020ff232b6551fa8bdd/ykdl-1.8.2.zip"
	ykdl_cmd_zip := exec.Command("curl", "-L", ykdl_zip_url, "-o", "package.zip")
	err = ykdl_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	ykdl_bin_url := "https://files.pythonhosted.org/packages/f2/27/f4e7616a139c84a04edb7778db2b3cfb77348ab73020ff232b6551fa8bdd/ykdl-1.8.2.bin"
	ykdl_cmd_bin := exec.Command("curl", "-L", ykdl_bin_url, "-o", "binary.bin")
	err = ykdl_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	ykdl_src_url := "https://files.pythonhosted.org/packages/f2/27/f4e7616a139c84a04edb7778db2b3cfb77348ab73020ff232b6551fa8bdd/ykdl-1.8.2.src.tar.gz"
	ykdl_cmd_src := exec.Command("curl", "-L", ykdl_src_url, "-o", "source.tar.gz")
	err = ykdl_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	ykdl_cmd_direct := exec.Command("./binary")
	err = ykdl_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: python@3.12")
exec.Command("latte", "install", "python@3.12")
}
