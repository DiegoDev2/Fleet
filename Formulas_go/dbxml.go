package main

// Dbxml - Embeddable XML database with XQuery support and other advanced features
// Homepage: https://www.oracle.com/database/technologies/related/berkeleydb.html

import (
	"fmt"
	
	"os/exec"
)

func installDbxml() {
	// Método 1: Descargar y extraer .tar.gz
	dbxml_tar_url := "https://download.oracle.com/berkeley-db/dbxml-6.1.4.tar.gz"
	dbxml_cmd_tar := exec.Command("curl", "-L", dbxml_tar_url, "-o", "package.tar.gz")
	err := dbxml_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	dbxml_zip_url := "https://download.oracle.com/berkeley-db/dbxml-6.1.4.zip"
	dbxml_cmd_zip := exec.Command("curl", "-L", dbxml_zip_url, "-o", "package.zip")
	err = dbxml_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	dbxml_bin_url := "https://download.oracle.com/berkeley-db/dbxml-6.1.4.bin"
	dbxml_cmd_bin := exec.Command("curl", "-L", dbxml_bin_url, "-o", "binary.bin")
	err = dbxml_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	dbxml_src_url := "https://download.oracle.com/berkeley-db/dbxml-6.1.4.src.tar.gz"
	dbxml_cmd_src := exec.Command("curl", "-L", dbxml_src_url, "-o", "source.tar.gz")
	err = dbxml_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	dbxml_cmd_direct := exec.Command("./binary")
	err = dbxml_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: berkeley-db")
exec.Command("latte", "install", "berkeley-db")
	fmt.Println("Instalando dependencia: xerces-c")
exec.Command("latte", "install", "xerces-c")
	fmt.Println("Instalando dependencia: xqilla")
exec.Command("latte", "install", "xqilla")
}
