import {
  Button,
  FormControl,
  FormErrorMessage,
  FormLabel,
  Input,
  Modal,
  ModalBody,
  ModalCloseButton,
  ModalContent,
  ModalFooter,
  ModalHeader,
  ModalOverlay,
} from "@chakra-ui/react"
import { useMutation, useQueryClient } from "@tanstack/react-query"
import { type SubmitHandler, useForm } from "react-hook-form"

import {
  type ApiError,
  type BookmarkPublic,
  type BookmarkUpdate,
  BookmarksService,
} from "../../client"
import useCustomToast from "../../hooks/useCustomToast"
import { handleError } from "../../utils"

interface EditBookmarkProps {
  bookmark: BookmarkPublic
  isOpen: boolean
  onClose: () => void
}

const EditBookmark = ({ bookmark, isOpen, onClose }: EditBookmarkProps) => {
  const queryClient = useQueryClient()
  const showToast = useCustomToast()
  const {
    register,
    handleSubmit,
    reset,
    formState: { isSubmitting, errors, isDirty },
  } = useForm<BookmarkUpdate>({
    mode: "onBlur",
    criteriaMode: "all",
    defaultValues: bookmark,
  })

  const mutation = useMutation({
    mutationFn: (data: BookmarkUpdate) =>
      BookmarksService.updateBookmark({ id: bookmark.id, requestBody: data }),
    onSuccess: () => {
      showToast("Success!", "Bookmark updated successfully.", "success")
      onClose()
    },
    onError: (err: ApiError) => {
      handleError(err, showToast)
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ["items"] })
    },
  })

  const onSubmit: SubmitHandler<BookmarkUpdate> = async (data) => {
    mutation.mutate(data)
  }

  const onCancel = () => {
    reset()
    onClose()
  }

  return (
    <>
      <Modal
        isOpen={isOpen}
        onClose={onClose}
        size={{ base: "sm", md: "md" }}
        isCentered
      >
        <ModalOverlay />
        <ModalContent as="form" onSubmit={handleSubmit(onSubmit)}>
          <ModalHeader>Edit Bookmark</ModalHeader>
          <ModalCloseButton />
          <ModalBody pb={6}>
            <FormControl isInvalid={!!errors.url}>
              <FormLabel htmlFor="url">URL</FormLabel>
              <Input
                id="url"
                {...register("url", {
                  required: "URL is required",
                })}
                type="text"
              />
              {errors.url && (
                <FormErrorMessage>{errors.url.message}</FormErrorMessage>
              )}
            </FormControl>

            <FormControl isInvalid={!!errors.title}>
              <FormLabel htmlFor="title">Title</FormLabel>
              <Input
                id="title"
                {...register("title", {
                  required: "Title is required",
                })}
                type="text"
              />
              {errors.title && (
                <FormErrorMessage>{errors.title.message}</FormErrorMessage>
              )}
            </FormControl>
            <FormControl mt={4}>
              <FormLabel htmlFor="description">Description</FormLabel>
              <Input
                id="description"
                {...register("description")}
                placeholder="Description"
                type="text"
              />
            </FormControl>
            <FormControl mt={4}>
              <FormLabel htmlFor="tags">Tags</FormLabel>
              <Input
                id="tags"
                {...register("tags")}
                placeholder="Tags"
                type="text"
              />
            </FormControl>

          </ModalBody>
          <ModalFooter gap={3}>
            <Button
              variant="primary"
              type="submit"
              isLoading={isSubmitting}
              isDisabled={!isDirty}
            >
              Save
            </Button>
            <Button onClick={onCancel}>Cancel</Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  )
}

export default EditBookmark
