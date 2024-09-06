package main

// Crm114 - Examine, sort, filter or alter logs or data streams
// Homepage: https://crm114.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installCrm114() {
	// Método 1: Descargar y extraer .tar.gz
	crm114_tar_url := "https://crm114.sourceforge.net/tarballs/crm114-20100106-BlameMichelson.src.tar.gz"
	crm114_cmd_tar := exec.Command("curl", "-L", crm114_tar_url, "-o", "package.tar.gz")
	err := crm114_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	crm114_zip_url := "https://crm114.sourceforge.net/tarballs/crm114-20100106-BlameMichelson.src.zip"
	crm114_cmd_zip := exec.Command("curl", "-L", crm114_zip_url, "-o", "package.zip")
	err = crm114_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	crm114_bin_url := "https://crm114.sourceforge.net/tarballs/crm114-20100106-BlameMichelson.src.bin"
	crm114_cmd_bin := exec.Command("curl", "-L", crm114_bin_url, "-o", "binary.bin")
	err = crm114_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	crm114_src_url := "https://crm114.sourceforge.net/tarballs/crm114-20100106-BlameMichelson.src.src.tar.gz"
	crm114_cmd_src := exec.Command("curl", "-L", crm114_src_url, "-o", "source.tar.gz")
	err = crm114_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	crm114_cmd_direct := exec.Command("./binary")
	err = crm114_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: tre")
exec.Command("latte", "install", "tre")
}
