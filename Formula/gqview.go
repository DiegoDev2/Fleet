package main

// Gqview - Image browser
// Homepage: https://gqview.sourceforge.net/

import (
	"fmt"
	
	"os/exec"
)

func installGqview() {
	// Método 1: Descargar y extraer .tar.gz
	gqview_tar_url := "https://downloads.sourceforge.net/project/gqview/gqview/2.0.4/gqview-2.0.4.tar.gz"
	gqview_cmd_tar := exec.Command("curl", "-L", gqview_tar_url, "-o", "package.tar.gz")
	err := gqview_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gqview_zip_url := "https://downloads.sourceforge.net/project/gqview/gqview/2.0.4/gqview-2.0.4.zip"
	gqview_cmd_zip := exec.Command("curl", "-L", gqview_zip_url, "-o", "package.zip")
	err = gqview_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gqview_bin_url := "https://downloads.sourceforge.net/project/gqview/gqview/2.0.4/gqview-2.0.4.bin"
	gqview_cmd_bin := exec.Command("curl", "-L", gqview_bin_url, "-o", "binary.bin")
	err = gqview_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gqview_src_url := "https://downloads.sourceforge.net/project/gqview/gqview/2.0.4/gqview-2.0.4.src.tar.gz"
	gqview_cmd_src := exec.Command("curl", "-L", gqview_src_url, "-o", "source.tar.gz")
	err = gqview_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gqview_cmd_direct := exec.Command("./binary")
	err = gqview_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: gtk+")
	exec.Command("latte", "install", "gtk+").Run()
}
