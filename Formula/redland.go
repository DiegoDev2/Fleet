package main

// Redland - RDF Library
// Homepage: https://librdf.org/

import (
	"fmt"
	
	"os/exec"
)

func installRedland() {
	// Método 1: Descargar y extraer .tar.gz
	redland_tar_url := "https://download.librdf.org/source/redland-1.0.17.tar.gz"
	redland_cmd_tar := exec.Command("curl", "-L", redland_tar_url, "-o", "package.tar.gz")
	err := redland_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	redland_zip_url := "https://download.librdf.org/source/redland-1.0.17.zip"
	redland_cmd_zip := exec.Command("curl", "-L", redland_zip_url, "-o", "package.zip")
	err = redland_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	redland_bin_url := "https://download.librdf.org/source/redland-1.0.17.bin"
	redland_cmd_bin := exec.Command("curl", "-L", redland_bin_url, "-o", "binary.bin")
	err = redland_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	redland_src_url := "https://download.librdf.org/source/redland-1.0.17.src.tar.gz"
	redland_cmd_src := exec.Command("curl", "-L", redland_src_url, "-o", "source.tar.gz")
	err = redland_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	redland_cmd_direct := exec.Command("./binary")
	err = redland_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: raptor")
	exec.Command("latte", "install", "raptor").Run()
	fmt.Println("Instalando dependencia: rasqal")
	exec.Command("latte", "install", "rasqal").Run()
	fmt.Println("Instalando dependencia: sqlite")
	exec.Command("latte", "install", "sqlite").Run()
	fmt.Println("Instalando dependencia: unixodbc")
	exec.Command("latte", "install", "unixodbc").Run()
}
