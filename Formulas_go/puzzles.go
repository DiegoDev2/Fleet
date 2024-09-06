package main

// Puzzles - Collection of one-player puzzle games
// Homepage: https://www.chiark.greenend.org.uk/~sgtatham/puzzles/

import (
	"fmt"
	
	"os/exec"
)

func installPuzzles() {
	// Método 1: Descargar y extraer .tar.gz
	puzzles_tar_url := "https://www.chiark.greenend.org.uk/~sgtatham/puzzles/puzzles-20240827.52afffa.tar.gz"
	puzzles_cmd_tar := exec.Command("curl", "-L", puzzles_tar_url, "-o", "package.tar.gz")
	err := puzzles_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	puzzles_zip_url := "https://www.chiark.greenend.org.uk/~sgtatham/puzzles/puzzles-20240827.52afffa.zip"
	puzzles_cmd_zip := exec.Command("curl", "-L", puzzles_zip_url, "-o", "package.zip")
	err = puzzles_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	puzzles_bin_url := "https://www.chiark.greenend.org.uk/~sgtatham/puzzles/puzzles-20240827.52afffa.bin"
	puzzles_cmd_bin := exec.Command("curl", "-L", puzzles_bin_url, "-o", "binary.bin")
	err = puzzles_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	puzzles_src_url := "https://www.chiark.greenend.org.uk/~sgtatham/puzzles/puzzles-20240827.52afffa.src.tar.gz"
	puzzles_cmd_src := exec.Command("curl", "-L", puzzles_src_url, "-o", "source.tar.gz")
	err = puzzles_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	puzzles_cmd_direct := exec.Command("./binary")
	err = puzzles_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cmake")
exec.Command("latte", "install", "cmake")
	fmt.Println("Instalando dependencia: halibut")
exec.Command("latte", "install", "halibut")
	fmt.Println("Instalando dependencia: imagemagick")
exec.Command("latte", "install", "imagemagick")
	fmt.Println("Instalando dependencia: pkg-config")
exec.Command("latte", "install", "pkg-config")
	fmt.Println("Instalando dependencia: cairo")
exec.Command("latte", "install", "cairo")
	fmt.Println("Instalando dependencia: gdk-pixbuf")
exec.Command("latte", "install", "gdk-pixbuf")
	fmt.Println("Instalando dependencia: glib")
exec.Command("latte", "install", "glib")
	fmt.Println("Instalando dependencia: gtk+3")
exec.Command("latte", "install", "gtk+3")
	fmt.Println("Instalando dependencia: pango")
exec.Command("latte", "install", "pango")
}
