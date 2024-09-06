package main

// Convertlit - Convert Microsoft Reader format eBooks into open format
// Homepage: http://www.convertlit.com/

import (
	"fmt"
	
	"os/exec"
)

func installConvertlit() {
	// Método 1: Descargar y extraer .tar.gz
	convertlit_tar_url := "http://www.convertlit.com/clit18src.zip"
	convertlit_cmd_tar := exec.Command("curl", "-L", convertlit_tar_url, "-o", "package.tar.gz")
	err := convertlit_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	convertlit_zip_url := "http://www.convertlit.com/clit18src.zip"
	convertlit_cmd_zip := exec.Command("curl", "-L", convertlit_zip_url, "-o", "package.zip")
	err = convertlit_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	convertlit_bin_url := "http://www.convertlit.com/clit18src.zip"
	convertlit_cmd_bin := exec.Command("curl", "-L", convertlit_bin_url, "-o", "binary.bin")
	err = convertlit_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	convertlit_src_url := "http://www.convertlit.com/clit18src.zip"
	convertlit_cmd_src := exec.Command("curl", "-L", convertlit_src_url, "-o", "source.tar.gz")
	err = convertlit_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	convertlit_cmd_direct := exec.Command("./binary")
	err = convertlit_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: libtommath")
	exec.Command("latte", "install", "libtommath").Run()
}
