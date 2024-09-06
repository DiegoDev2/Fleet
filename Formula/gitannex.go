package main

// GitAnnex - Manage files with git without checking in file contents
// Homepage: https://git-annex.branchable.com/

import (
	"fmt"
	
	"os/exec"
)

func installGitAnnex() {
	// Método 1: Descargar y extraer .tar.gz
	gitannex_tar_url := "https://hackage.haskell.org/package/git-annex-10.20240831/git-annex-10.20240831.tar.gz"
	gitannex_cmd_tar := exec.Command("curl", "-L", gitannex_tar_url, "-o", "package.tar.gz")
	err := gitannex_cmd_tar.Run()
	if err != nil {
		fmt.Println("Error al descargar .tar.gz:", err)
		return
	}
	exec.Command("tar", "-xzf", "package.tar.gz").Run()

	// Método 2: Descargar y extraer .zip
	gitannex_zip_url := "https://hackage.haskell.org/package/git-annex-10.20240831/git-annex-10.20240831.zip"
	gitannex_cmd_zip := exec.Command("curl", "-L", gitannex_zip_url, "-o", "package.zip")
	err = gitannex_cmd_zip.Run()
	if err != nil {
		fmt.Println("Error al descargar .zip:", err)
		return
	}
	exec.Command("unzip", "package.zip").Run()

	// Método 3: Descargar binario precompilado
	gitannex_bin_url := "https://hackage.haskell.org/package/git-annex-10.20240831/git-annex-10.20240831.bin"
	gitannex_cmd_bin := exec.Command("curl", "-L", gitannex_bin_url, "-o", "binary.bin")
	err = gitannex_cmd_bin.Run()
	if err != nil {
		fmt.Println("Error al descargar binario:", err)
		return
	}
	exec.Command("chmod", "+x", "binary.bin").Run()
	exec.Command("./binary.bin").Run()

	// Método 4: Descargar y compilar desde código fuente
	gitannex_src_url := "https://hackage.haskell.org/package/git-annex-10.20240831/git-annex-10.20240831.src.tar.gz"
	gitannex_cmd_src := exec.Command("curl", "-L", gitannex_src_url, "-o", "source.tar.gz")
	err = gitannex_cmd_src.Run()
	if err != nil {
		fmt.Println("Error al descargar código fuente:", err)
		return
	}
	exec.Command("tar", "-xzf", "source.tar.gz").Run()
	exec.Command("make").Run()

	// Método 5: Ejecutar binario directo
	gitannex_cmd_direct := exec.Command("./binary")
	err = gitannex_cmd_direct.Run()
	if err != nil {
		fmt.Println("Error al ejecutar binario:", err)
		return
	}
	// Instalar dependencias
	fmt.Println("Instalando dependencia: cabal-install")
	exec.Command("latte", "install", "cabal-install").Run()
	fmt.Println("Instalando dependencia: ghc@9.8")
	exec.Command("latte", "install", "ghc@9.8").Run()
	fmt.Println("Instalando dependencia: pkg-config")
	exec.Command("latte", "install", "pkg-config").Run()
	fmt.Println("Instalando dependencia: libmagic")
	exec.Command("latte", "install", "libmagic").Run()
}
