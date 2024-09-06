package main

// AntContrib - Collection of tasks for Apache Ant
// Homepage: https://ant-contrib.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installAntContrib() {
	// Método 1: Descargar y extraer .tar.gz
	antcontrib_tar_url := "https://downloads.sourceforge.net/project/ant-contrib/ant-contrib/1.0b3/ant-contrib-1.0b3-bin.tar.gz"
	antcontrib_cmd_tar := exec.Command("curl", "-L", antcontrib_tar_url, "-o", "package.tar.gz")
	err := antcontrib_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	antcontrib_zip_url := "https://downloads.sourceforge.net/project/ant-contrib/ant-contrib/1.0b3/ant-contrib-1.0b3-bin.zip"
	antcontrib_cmd_zip := exec.Command("curl", "-L", antcontrib_zip_url, "-o", "package.zip")
	err = antcontrib_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	antcontrib_bin_url := "https://downloads.sourceforge.net/project/ant-contrib/ant-contrib/1.0b3/ant-contrib-1.0b3-bin.bin"
	antcontrib_cmd_bin := exec.Command("curl", "-L", antcontrib_bin_url, "-o", "binary.bin")
	err = antcontrib_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	antcontrib_src_url := "https://downloads.sourceforge.net/project/ant-contrib/ant-contrib/1.0b3/ant-contrib-1.0b3-bin.src.tar.gz"
	antcontrib_cmd_src := exec.Command("curl", "-L", antcontrib_src_url, "-o", "source.tar.gz")
	err = antcontrib_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	antcontrib_cmd_direct := exec.Command("./binary")
	err = antcontrib_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: ant")
exec.Command("latte", "install", "ant")
}
