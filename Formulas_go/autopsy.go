package main

// Autopsy - Graphical interface to Sleuth Kit investigation tools
// Homepage: https://www.sleuthkit.org/autopsy/index.php

import (
	"fmt"
	
	"os/exec"
)

func installAutopsy() {
	// Método 1: Descargar y extraer .tar.gz
	autopsy_tar_url := "https://downloads.sourceforge.net/project/autopsy/autopsy/2.24/autopsy-2.24.tar.gz"
	autopsy_cmd_tar := exec.Command("curl", "-L", autopsy_tar_url, "-o", "package.tar.gz")
	err := autopsy_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	autopsy_zip_url := "https://downloads.sourceforge.net/project/autopsy/autopsy/2.24/autopsy-2.24.zip"
	autopsy_cmd_zip := exec.Command("curl", "-L", autopsy_zip_url, "-o", "package.zip")
	err = autopsy_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	autopsy_bin_url := "https://downloads.sourceforge.net/project/autopsy/autopsy/2.24/autopsy-2.24.bin"
	autopsy_cmd_bin := exec.Command("curl", "-L", autopsy_bin_url, "-o", "binary.bin")
	err = autopsy_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	autopsy_src_url := "https://downloads.sourceforge.net/project/autopsy/autopsy/2.24/autopsy-2.24.src.tar.gz"
	autopsy_cmd_src := exec.Command("curl", "-L", autopsy_src_url, "-o", "source.tar.gz")
	err = autopsy_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	autopsy_cmd_direct := exec.Command("./binary")
	err = autopsy_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: sleuthkit")
exec.Command("latte", "install", "sleuthkit")
	fmt.Println("Instalando dependencia: file-formula")
exec.Command("latte", "install", "file-formula")
	fmt.Println("Instalando dependencia: grep")
exec.Command("latte", "install", "grep")
	fmt.Println("Instalando dependencia: md5sha1sum")
exec.Command("latte", "install", "md5sha1sum")
}
