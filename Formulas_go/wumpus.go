package main

// Wumpus - Exact clone of the ancient BASIC Hunt the Wumpus game
// Homepage: http://www.catb.org/~esr/wumpus/

import (
	"fmt"
	
	"os/exec"
)

func installWumpus() {
	// Método 1: Descargar y extraer .tar.gz
	wumpus_tar_url := "http://www.catb.org/~esr/wumpus/wumpus-1.10.tar.gz"
	wumpus_cmd_tar := exec.Command("curl", "-L", wumpus_tar_url, "-o", "package.tar.gz")
	err := wumpus_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	wumpus_zip_url := "http://www.catb.org/~esr/wumpus/wumpus-1.10.zip"
	wumpus_cmd_zip := exec.Command("curl", "-L", wumpus_zip_url, "-o", "package.zip")
	err = wumpus_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	wumpus_bin_url := "http://www.catb.org/~esr/wumpus/wumpus-1.10.bin"
	wumpus_cmd_bin := exec.Command("curl", "-L", wumpus_bin_url, "-o", "binary.bin")
	err = wumpus_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	wumpus_src_url := "http://www.catb.org/~esr/wumpus/wumpus-1.10.src.tar.gz"
	wumpus_cmd_src := exec.Command("curl", "-L", wumpus_src_url, "-o", "source.tar.gz")
	err = wumpus_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	wumpus_cmd_direct := exec.Command("./binary")
	err = wumpus_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
}
