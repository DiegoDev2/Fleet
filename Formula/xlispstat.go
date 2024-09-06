package main

// Xlispstat - Statistical data science environment based on Lisp
// Homepage: https://homepage.stat.uiowa.edu/~luke/xls/xlsinfo/

import (
	"fmt"
	
	"os/exec"
)

func installXlispstat() {
	// Método 1: Descargar y extraer .tar.gz
	xlispstat_tar_url := "https://homepage.cs.uiowa.edu/~luke/xls/xlispstat/current/xlispstat-3-52-23.tar.gz"
	xlispstat_cmd_tar := exec.Command("curl", "-L", xlispstat_tar_url, "-o", "package.tar.gz")
	err := xlispstat_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	xlispstat_zip_url := "https://homepage.cs.uiowa.edu/~luke/xls/xlispstat/current/xlispstat-3-52-23.zip"
	xlispstat_cmd_zip := exec.Command("curl", "-L", xlispstat_zip_url, "-o", "package.zip")
	err = xlispstat_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	xlispstat_bin_url := "https://homepage.cs.uiowa.edu/~luke/xls/xlispstat/current/xlispstat-3-52-23.bin"
	xlispstat_cmd_bin := exec.Command("curl", "-L", xlispstat_bin_url, "-o", "binary.bin")
	err = xlispstat_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	xlispstat_src_url := "https://homepage.cs.uiowa.edu/~luke/xls/xlispstat/current/xlispstat-3-52-23.src.tar.gz"
	xlispstat_cmd_src := exec.Command("curl", "-L", xlispstat_src_url, "-o", "source.tar.gz")
	err = xlispstat_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	xlispstat_cmd_direct := exec.Command("./binary")
	err = xlispstat_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libx11")
	exec.Command("latte", "install", "libx11").Run()
}
